# Random Oracle
**Oracle for sourcing on-chain randomness**


#### Local Testnet
```
make start-localnet
```

#### Developing

**Generate Protos**
```shell
make proto-gen
```

**Generate Mocks**

```shell
mockgen --source=./provider/expected_provider.go --destination=./testutil/expected_keepers_mocks.go --package=testutil

```