const Web3 = require('web3')
const detectNetwork = require('web3-detect-network')
const web3 = new Web3()
const BN = require('bn.js')

const big = (n) => new BN(n.toString(10))

const contractJSON = require('../ethereum-contracts/build/contracts/Exchange.json')
const { abi, networks } = contractJSON
const { address } = networks[Object.keys(networks)[0]]

function getWeb3() {
  return web3
}

function toWei(value) {
  return web3.utils.toWei(`${value||0}`, 'ether')
}

function toEth(value) {
  return web3.utils.fromWei(`${value||0}`, 'ether')
}

function getModels() {
  return window.loopbackApp.models
}

function getContractAddress() {
  return address
}

function getContractAbi() {
  return abi
}

function getWindowWeb3() {
  return window.web3
}

function isMetamaskConnected() {
  const web3 = getWindowWeb3()
  if (!web3) return false
  return web3.currentProvider && web3.currentProvider && web3.currentProvider.isConnected()
}

function getProvider() {
  const defaultProvider = web3.providers.HttpProvider('http://localhost:8545')
  const w3 = getWindowWeb3()
  if (!w3) return defaultProvider
  return w3.currentProvider
}

function getConnectedAccount() {
  return isMetamaskConnected() && getWindowWeb3().eth.defaultAccount
}

async function getConnectedNetwork() {
  return detectNetwork(getProvider())
}

async function getBalance(account) {
  return new Promise((resolve, reject) => {
    getWindowWeb3().eth.getBalance(account, (err, result) => {
      if (err) return reject(err)
      resolve(result)
    })
  })
}

async function placeOrder(opts) {
  const { amount, price, value, account } = opts
  const web3 = new Web3(getProvider())
  const instance = new web3.eth.Contract(abi, address)
  const result = await instance.methods.placeOrder(amount, price).send({
    from: account,
    value
  })

  return result
}

async function getContractBalance(account) {
  const web3 = new Web3(getProvider())
  const instance = new web3.eth.Contract(abi, address)
  const result = await instance.methods.deposits(account).call()

  return big(result.amount)
}

async function main() {
  const sellForm = document.querySelector("#sellForm")
  const buyAmount = document.querySelector("#buyAmount")
  const buyPrice = document.querySelector("#buyPrice")
  const depositAmount = document.querySelector("#depositAmount")

  sellForm.addEventListener("submit", async function(event) {
    event.preventDefault()

    console.log(getConnectedAccount())

    try {
      const result = await placeOrder({
        amount: toWei(buyAmount.value),
        price: toWei(buyPrice.value),
        value: toWei(depositAmount.value),
        account: getConnectedAccount()
      })

      console.log(result)
    } catch(err) {
      alert(err)
    }
  })
}

main()
