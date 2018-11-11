const contractJSON = require('../build/contracts/Exchange.json')
const { abi, networks } = contractJSON
const { address } = networks[Object.keys(networks)[0]]

console.log(address)
