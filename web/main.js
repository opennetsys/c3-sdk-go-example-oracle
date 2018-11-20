const Web3 = require('web3')
const detectNetwork = require('web3-detect-network')
const web3 = new Web3()
const BN = require('bn.js')
const ScatterJS = require('scatterjs-core').default
const ScatterEOS = require('scatterjs-plugin-eosjs').default
const Eos = require('eosjs')

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


  const buyForm = document.querySelector("#buyForm")
  const sellAmount = document.querySelector("#sellAmount")
  const sellPrice = document.querySelector("#sellPrice")
  const depositAmountEos = document.querySelector("#depositAmountEos")


  ScatterJS.plugins(new ScatterEOS())
  const connectionOptions = {initTimeout:10000}

  const network = {
      blockchain: 'eos',
      protocol: 'https',
      host: 'api-kylin.eosasia.one',
      port: 443,
      chainId: '5fff1dae8dc8e2fc4d5b23b2c7665c97f9e9d8edf2b6485a86ba311c25639191'
  }

  async function placeOrderEos(opts) {
    const { amount, price, value } = opts

    const scatter = ScatterJS.scatter;
    const requiredFields = { accounts:[network] };
    scatter.getIdentity(requiredFields).then(async () => {
      const account = scatter.identity.accounts.find(x => x.blockchain === 'eos');
      const eosOptions = { expireInSeconds:60 }
      // Get a proxy reference to eosjs which you can use to sign transactions with a user's Scatter.
      const eos = scatter.eos(network, Eos, eosOptions);
      const transactionOptions = { authorization:[`${account.name}@${account.authority}`] };

      const result = await eos.transaction({
        actions: [{
          account: 'helloworld54',
          name: 'placeorder',
          authorization: [{
            actor: 'myaccount123',
            permission: 'active',
          }],
          data: {
            acct: account.name,
            price: Number(price),
            amount: Number(amount),
            value: Number(value),
          },
        }]
      }, {
        broadcast: true,
        sign: true
      });

      console.log(result.transaction_id)

      /*
      eos.transfer(account.name, 'helloworld54', `${amount}.0000 EOS`, 'memo', transactionOptions).then(trx => {
      console.log(`Transaction ID: ${trx.transaction_id}`);
      }).catch(error => {
      console.error(error);
      });
      */
    })
    .catch(err => {
      console.error(err)
    })
  }

  ScatterJS.scatter.connect("My-App", connectionOptions).then(connected => {
    console.log('connected', connected)
      if (!connected) {
          alert('scatter is not installed')
          return false;
      }
  });

  sellForm.addEventListener("submit", async function(event) {
    event.preventDefault()

    console.log(getConnectedAccount())

    try {
      const result = await placeOrder({
        amount: toWei(buyAmount.value),
        price: toWei(buyPrice.value),
        value: toWei(depositAmount.value)
      })

      console.log(result)
    } catch(err) {
      alert(err)
    }
  })

  buyForm.addEventListener("submit", async function(event) {
    event.preventDefault()

    try {
      placeOrderEos({
        amount: sellAmount.value,
        price: sellPrice.value,
        value: depositAmountEos.value,
      })
    } catch(err) {
      alert(err)
    }
  })
}

main()
