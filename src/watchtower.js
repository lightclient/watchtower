const axios = require('axios')
const uuidv4 = require('uuid/v4')

const twilio = require('./clients/twilio')
const twitter = require('./clients/twitter')
const redis = require('./clients/redis')

const utils = require('./utils')

// const defaultOptions = {
//   port: '9898',
//   dbProvider: LevelDBProvider,
//   operatorProvider: PlasmaCore.providers.OperatorProviders.HttpOperatorProvider,
//   contractProvider: PlasmaCore.providers.ContractProviders.ContractProvider,
//   walletProvider: PlasmaCore.providers.WalletProviders.LocalWalletProvider,
//   dbPath: dbPaths.CHAIN_DB_PATH
// }

/**
 * Class that houses the Watchtower.
 */
class Watchtower {
  constructor (options) {
    this.options = Object.assign({}, options)
    console.log(options)
  }

  /**
   * Starts the node.
   */
  async start () {
    this.started = true
    await this._initConnection()
    await this._sendRPC('pg_monitorAccount', this.options.address)
    this._pollNode()
  }

  /**
   * Stops the node.
   */
  stop () {
    this.started = false
  }

  /**
   * Polls the node for invalid transactions and notifies the account owner if
   * someone attempts to transfer a coinrange they own
   */
  async _pollNode() {
    if (!this.started) return

    try {
      console.log('polling')

      const response = await this._sendRPC(
        'pg_getMaliciousTransactions', 
        this.options.address
      )

      console.log('got this from node: ', response)

      for (element in response.result) {
        const [txHash, error] = element
        
        const notified = await redis.get(txHash)

        if (!notified) {
          redis.set(txHash, 'true')
          this._notifyOwner()
        }
      }
    } finally {
      await utils.sleep(10000)
      this._pollOperator()
    }
  }

  /**
   * Sends the owner a text message regarding the invalid transfer attempt and
   * posts a tweet to the Plasma chain's Twitter account
   */
  async _notifyOwner() {
    await Promise.all([
      this._sendSMS(),
      this._sendTweet(),
    ])
  }

  /**
   * Sends a SMS message to the account owner
   */
  async _sendSMS(data) {
    console.log('sending sms')
    try {
      await twilio.send('Your coins are running away!', this.options.phone)
    } catch (e) {
      console.log(e)
    }
  }

  /**
   * Post a Tweet to the Plasma chain's Twitter account
   */
  async _sendTweet(data) {
    console.log('sending tweet')
    try {
      await twitter.post('Bad stuff is happening on the chain!')
    } catch(e) {
      console.log(e)
    }
  }

  /**
   * Initializes axios http connection to the Plasma node
   */
  async _initConnection () {
    this.endpoint = 'http://plasma-node-cluster-ip-service:9898'
    this.http = axios.create({
      baseURL: this.endpoint.startsWith('http')
        ? this.endpoint
        : `https://${this.endpoint}`
    })
  }

  /**
   * Sends an RPC call over http to the Plasma node
   * @param {string} method 
   * @param {string | Object} params 
   */
  async _sendRPC(method, params) {
    try {
      const response = await this.http.post('/', {
        jsonrpc: '2.0',
          method: method,
          params: params,
          id: uuidv4()
      })

      return JSON.parse(response.data)
    } catch(e) {
      throw e
    }
  }
}

module.exports = Watchtower
