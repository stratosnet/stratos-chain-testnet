# stratos-chain-testnet
stratos block testnet genesis and config

## How to connect Stratos Chain Testnet

### Download release binary

#### 1. get stchaind & stchaincli binary file
```bash
wget https://github.com/stratosnet/stratos-chain/releases/download/v0.3.0/stchaincli
wget https://github.com/stratosnet/stratos-chain/releases/download/v0.3.0/stchaind
```
then check the granularity 
```bash
md5sum *

## expect output 
## 604200d7c85802fac02029f474d5ae1e  stchaincli
## ede9385b6ba4d4b5101ae06b33f4a2e9  stchaind
```

add execute permission for the binary downloaded
```bash
chmod +x stchaincli
chmod +x stchaind
```

#### 2. get the genesis and config file
initialize the node
```bash
./stchaind init --home ./  "<node mae you prefer>"

# ignore the output since you need to download the genesis file 
```

```bash
wget https://raw.githubusercontent.com/stratosnet/stratos-chain-testnet/main/genesis.json
wget https://raw.githubusercontent.com/stratosnet/stratos-chain-testnet/main/config.toml
```

change your node moniker in config.toml (optional if you don't want to become validator)
```bash
# A custom human readable name for this node
moniker = "<node name you prefer>"
```

move the config.toml and genesis.json file to config folder
```bash
mv config.toml config/
mv genesis.json config/
```
#### 3. run the node

```bash
./stchaind start --home ./ 
```
after this, the node will try to catch up with the blockchain to the latest block

you can run the node in background
```bash
./stchaind start --home ./ 2>&1 >> chain.log &
```


#### for more info about get test token from faucet and send tx. 
(TBA)
