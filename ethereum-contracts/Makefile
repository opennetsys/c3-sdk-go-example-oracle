
.PHONY: build
build:
	@rm -rf build/ && truffle compile

.PHONY: deps
deps:
	@npm i
	@npm i -g truffle

.PHONY: abi
abi:
	@node scripts/abi.js > abi/Exchange.abi && abigen --abi=abi/Exchange.abi --pkg=exchange --out=../contracts/Exchange.go

.PHONY: deploy
deploy:
	@rm -rf build/ && truffle deploy

.PHONY: deploy/rinkeby
deploy/rinkeby:
	@rm -rf build/ && truffle deploy --network=rinkeby

.PHONY: start/testrpc
start/testrpc:
	@ganache-cli -m "benefit come fiscal degree photo you abstract casino heart habit early pact"

.PHONY: test
test:
	@truffle test

.PHONY: address
address:
	@node scripts/contract_address.js
