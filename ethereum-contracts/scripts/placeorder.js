require('dotenv').config()
const Web3 = require('web3')
const assert = require('assert')
const BN = require('bn.js')

const big = (n) => new BN(n.toString(10))
const tenPow18 = big(10).pow(big(18))
const toEth = n => big(n).mul(tenPow18)

const contract = require('../build/contracts/Exchange.json')
const abi = contract.abi
const contractAddress = contract.networks[Object.keys(contract.networks)[0]].address

const provider = new Web3.providers.HttpProvider('https://rinkeby.infura.io')
//const provider = new Web3.providers.HttpProvider('http://localhost:8545')
const web3 = new Web3(provider)

const instance = new web3.eth.Contract(abi, contractAddress)

async function main() {
  const accounts = await web3.eth.getAccounts()

  const hub = accounts[0]
  //const alice = accounts[1]
  const alice = '0x656f3db0b3a18a0e2c80f7d55f8eb9fd813e19c2'
  const bob = accounts[2]

  const amount = toEth(1).toString(10)
  const price = toEth(1).toString(10)
  const value = toEth(1).toString(10)

  const data = await instance.methods.placeOrder(amount, price).encodeABI()

  /*
  const result = await instance.methods.placeOrder(amount, price).send({
    from: alice,
    value: value,
    gas: 4500000,
    gasPrice: 10000000000,
  })
  */

  const gas = 4712383
  const gasPrice = 20000000000

  const tx = {
    to: contractAddress,
    from: alice,
    value: value.toString(),
    gas,
    gasPrice,
    data
  }

  const privateKey = process.env.PRIVATE_KEY
  const signedTx = await web3.eth.accounts.signTransaction(tx, privateKey)

  const bal = big(await web3.eth.getBalance(alice))
  console.log(bal.toString(10))

  const result = await web3.eth.sendSignedTransaction(signedTx.rawTransaction)

  assert.equal(result.status, true)
  console.log('done')
}

main()
