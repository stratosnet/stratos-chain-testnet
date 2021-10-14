# stratos-chain-testnet
stratos block testnet genesis and config

## How to connect Stratos Chain Testnet

### Prepare environment to run node

#### 0. Create user account(optional)
to create separated and more secure environment it is recommended to create separate user account to run node
```
sudo adduser stratos --home /home/stratos
```
once user is created login to system using *stratos* account and proceed with installation steps in context of that user

### Download release binary

#### 1. get stchaind & stchaincli binary file
PLEASE NOTE: these binary are built using linux amd64, so if you are preparing run a node on different kernel, please follow step 1.1 to build binary for your machine

```bash
wget https://github.com/stratosnet/stratos-chain/releases/download/v0.5.0/stchaincli
wget https://github.com/stratosnet/stratos-chain/releases/download/v0.5.0/stchaind
```
then check the granularity 
```bash
md5sum stchain*

## expect output 
## aded93fa2f1f4816375dcb425ec4ce42 stchaincli
## 2d79092bd2096baa82879edc64c17294 stchaind
```

add execute permission for the binary downloaded
```bash
chmod +x stchaincli
chmod +x stchaind
```

#### 1.1 compile the binary with source code
Make sure you have Go 1.15+ installed ([link](https://golang.org/doc/install)). 

```bash
git clone https://github.com/stratosnet/stratos-chain.git
cd stratos-chain
git checkout v0.5.0
make build
```
The binary can be found in ./build folder

#### 2. get the genesis and config file
initialize the node
```bash
./stchaind init --home ./  "<node name you prefer>"

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

#### create an account
```bash
./stchaincli keys add --hd-path "m/44'/606'/0'/0/0" --keyring-backend test --home ./ wallet1
```


#### Faucet 
faucet will be available at https://faucet-test.thestratos.org/
get some test token 
```bash
curl -X POST https://faucet-test.thestratos.org/faucet/<your address>

```
1 stos = 1000000000 ustos

the faucet is default to send 100 stos to the provided address

check balance (you need to wait for your node catching up with the network)
```bash

./stchaincli query account <your address> --home ./

```

to check node status
```bash
./stchaincli status --home ./
```


#### send your first tx

```bash
./stchaincli tx send <from address> <to address> <amount> --home ./ --keyring-backend test --chain-id test-chain-1 

# then input y for the pop up to confirm send
```

#### 4. run node as a service
**NOTE:** All below steps require *root* privileges

create file ```/lib/systemd/system/stratos.service``` with following content:
```
[Unit]
Description=Stratos Chain Node
After=network-online.target

[Service]
User=stratos
Group=stratos
ExecStart=/home/stratos/stchaind start --home=/home/stratos/.stchaind
Restart=on-failure
RestartSec=3
LimitNOFILE=8192

[Install]
WantedBy=multi-user.target
```
once service file is created you need to enable and start service:
```
systemctl daemon-reload
systemctl enable stratos.service
systemctl start stratos.service
```
to check if service is runnign as expected:
```
systemctl status stratos.service
```
to check the node log
```
journalctl -u stratos.service -f 
# exit with ctrl+c
```
