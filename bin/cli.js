#!/usr/bin/env node

require('dotenv').config()

const program = require('commander')
const colors = require('colors')
const Watchtower = require('../src/watchtower')

program
  .version('0.0.1')
  .option('-a, --address <address>', 'Public address to monitor')
  .option('-p, --phone <phone>', 'Phone number to recieve SMS notifications on')
  .parse(process.argv)

// const debug = program.debug ? 'debug:*' : ''

// const options = {
//   finalityDepth: program.finality,
//   port: program.port,
//   ethereumEndpoint: program.ethereum,
//   debug: `service:*,${debug}`,
//   contractProvider: PlasmaCore.providers.ContractProviders.ContractProvider,
//   walletProvider: wallets[program.wallet],
//   operatorProvider: PlasmaCore.providers.OperatorProviders.HttpOperatorProvider,
//   plasmaChainName: program.chain,
//   registryAddress: program.registry
// }

const options = {
  address: program.address,
  phone: program.phone,
}

const watchtower = new Watchtower(options)

// (async () => {
//   // const latest = await latestVersion(pkg.name)
//   // if (pkg.version !== latest) {
//   //   console.log(colors.red('ERROR: ') + 'Your plasma-client is out of date.')
//   //   console.log('Please update to the latest version by running:')
//   //   console.log(colors.green('npm install -g --upgrade plasma-client'))
//   //   console.log()
//   //   console.log(`You might also want to reset your database (this won't delete your accounts):`)
//   //   console.log(colors.green('plasma-cli killdb'))
//   //   return
//   // }



//   // console.log('Plasma Client v' + pkg.version + ' 🎉 🎉 🎉 ')

//   // console.log(getSectionTitle('DISCLAIMER'))
//   // console.log('Plasma Client is alpha software and will probably break.')
//   // console.log(`Please do NOT use this application with real money (unless you're willing to lose it).`)

//   // console.log(getSectionTitle('Available Accounts'))
//   // const accounts = await client.getAccounts()
//   // accounts.forEach((account, i) => {
//   //   const maxDigits = (accounts.length - 1).toString().length
//   //   const accountNumber = i.toString().padStart(maxDigits, '0')
//   //   console.log(`(${accountNumber}) ${account}`)
//   // })

//   // console.log(getSectionTitle('Client Information'))
//   // console.log(`Plasma Chain: ${program.chain}`)
//   // console.log(`Ethereum Node: ${program.ethereum}`)
//   // console.log(`Listening on: http://${program.hostname}:${program.port}`)

//   // console.log(getSectionTitle('Logs'))
// })()



async function main() {
  await watchtower.start()
}

main().then(() => console.log('done'))
