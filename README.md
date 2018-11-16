# c3-sdk-go-example-oracle

> Exampele of a cross-chain decentralized exchange without the use of a trusted third party.

This app uses a bridge chain in order to keep an order book of bids and asks of EOS and ETH.

## Usage

In terminal 1, start c3-go

```bash
c3-go
```

In terminal 2, start web interface

```bash
make web
```

In terminal 3, send genesis tx

```bash
make genesis
```

In terminal 3, start watcher

```bash
make start
```

## License

MIT
