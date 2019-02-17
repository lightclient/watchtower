const utils = require('./utils')
const twilio = require('./clients/twilio')

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
  }

  /**
   * Starts the node.
   */
  start () {
    this.started = true
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
      this._sendSMS('Your funds are running away!')
    } finally {
      // await utils.sleep(1000)
      // this._pollOperator()
    }
  }

  /**
   * Sends the owner a text message regarding the invalid transfer attempt and
   * posts a tweet to the Plasma chain's Twitter account
   */
  async _notifyOwner() {
    await Promise.all([
      this._sendSMS(),
      // this._sendTweet(),
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
  async _sendTweet() {
    throw 'not implemented'
  }
}

module.exports = Watchtower
