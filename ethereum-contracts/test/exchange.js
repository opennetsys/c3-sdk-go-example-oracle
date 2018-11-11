const Exchange = artifacts.require('Exchange')
const moment = require('moment')
const BN = require('bn.js')
const {soliditySha3: sha3} = require('web3-utils')
const util = require('ethereumjs-util')
const Reverter = require('./util/reverter')
const getLastEvent = require('./util/getLastEvent')

const big = n => new BN(n.toString(10))
const tenPow18 = big(10).pow(big(18))
const toEth = n => big(n).mul(tenPow18)

contract('Exchange', (accounts) => {
  const reverter = new Reverter(web3);
  const owner = accounts[0]
  const alice = accounts[1]
  const bob = accounts[2]
  let instance

  before('setup', async () => {
    instance = await Exchange.new(owner)

    await reverter.snapshot()
  })

  context('Exchange', () => {
    describe('[init]', () => {
      it('should crash if owner address is invalid', async () => {
        try {
          await Channel.new(0x0);
          assert.ok(false);
        } catch (e) { }
      })
      after(async () => {
        await reverter.revert()
      })
    })
    describe('[buy]', () => {
      it('should place buy order', async () => {
        const amount = toEth(1).toString(10)
        const price = toEth(1).toString(10)
        const value = toEth(1).toString(10)

        const result = await instance.buy(amount, price, {
          from: alice,
          value
        })

        assert.equal(result.receipt.status, '0x01')

        const eventLog = await getLastEvent(instance)
        assert.equal(eventLog.event, 'LogBuy')

        const bal = await web3.eth.getBalance(instance.address)
        assert.equal(bal.toString(), value)
      })
    })
  })
})
