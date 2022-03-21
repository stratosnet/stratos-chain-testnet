# stratos-chain-testnet
stratos block testnet genesis and config

## Prepare environment to run node

### 0. Create user account(optional)
To create a separated and more secure environment, it is recommended to create a separated user account to run node.
```
sudo adduser stratos --home /home/stratos
```
Once `user` is created, login the system using *stratos* account and proceed with installation steps in context of that user

### 1. Download release binary files

#### Get `stchaind` & `stchaincli` binary files

```bash
cd $HOME
wget https://github.com/stratosnet/stratos-chain/releases/download/v0.7.0/stchaincli
wget https://github.com/stratosnet/stratos-chain/releases/download/v0.7.0/stchaind
```

> These binary files are built using linux amd64, so if you are preparing run a node on different kernel, please follow step 1.1 to build binaries for your machine
>
> For ease of use, we recommend you save these files in your `$HOME` folder. In the follwoing, we suppose you are in the `$HOME` folder.

#### Check the granularity
```bash
md5sum stchain*

## expect output 
## f15f7bc91214582cd2d8ec58e2103ce9 stchaincli
## 296f5622062587f1a56a3dad681418c8 stchaind
```

#### Add `execute` permission for the binary downloaded
```bash
chmod +x stchaincli
chmod +x stchaind
```

### 1.1 Compile the binary with source code
#### Make sure you have Go 1.16+ installed ([link](https://golang.org/doc/install)).

```bash
git clone https://github.com/stratosnet/stratos-chain.git
cd stratos-chain
git checkout v0.7.0
make build
```
The binary files `stchaincli` and `stchaind` can be found in `build` folder. Then, move these two binary files to `$HOME`

```shell
mv build/stchaincli ./
mv build/stchaind ./
```

#### Install the binary files to `$GOPATH/bin`
```bash
make install
```

### 2. Get the `genesis` and `config` file
#### Initialize the node
```bash
./stchaind init "<node name you prefer>"

# ignore the output since you need to download the genesis file 
```

#### Download `genesis.json` and `config.toml` files
```bash
wget https://raw.githubusercontent.com/stratosnet/stratos-chain-testnet/main/genesis.json
wget https://raw.githubusercontent.com/stratosnet/stratos-chain-testnet/main/config.toml
```

#### Change `moniker`(optional if you don't want to become validator)
In `config.toml` file, at Line #16, there is a “moniker” field. Change it to any name you like. It’s your node name on the network.
```bash
# A custom human readable name for this node
moniker = "<node name you prefer>"
```

#### Move the `config.toml` and `genesis.json` files to `.stchaind/config/` folder
```bash
mv config.toml .stchaind/config/
mv genesis.json .stchaind/config/
```
> By default, the two binary executable files `stchaincli` and `stchaind` as well as the directory `.stchaind` have been saved or created in the `$HOME` folder. The `.stchaind` folder contains the node's configurations and data.

### 3. Run the node

```bash
./stchaind start 
```
After this, the node will try to catch up with the blockchain to the latest block.

You can run the node in background
```bash
./stchaind start 2>&1 >> chain.log &
```

## Operations
Once the node finishes catch-up, you can operate the node for various transactions(tx) and queries. You can find all the documents [here](https://stratos.gitbook.io/st-docs/).

In the following, we list some of commonly-used operations.
  
### Create an account

```bash
./stchaincli keys add --hd-path "m/44'/606'/0'/0/0" --keyring-backend test  <your wallet name>
```

Example
```bash
./stchaincli keys add --hd-path "m/44'/606'/0'/0/0" --keyring-backend test  wallet1
```
> After executed the above command, a `.stchaincli` will be created in your `$HOME` folder. The `.stchaincli` contains your wallet information with its address.

### `Faucet`
Faucet will be available at https://faucet-tropos.thestratos.org/ to get test tokens

```bash
curl -X POST https://faucet-tropos.thestratos.org/faucet/<your wallet address>
```

> * 1 stos = 1000000000 ustos
> * By default, faucet will send 100stos(100,000,000,000ustos) to the given wallet address
> * maximum 3 faucet requests to arbitrary wallet address from a single IP within an hour
> * maximum 1 faucet request to a fixed wallet address within an hour

Check balance (you need to wait for your node catching up with the network)
```bash
./stchaincli query account <your wallet address>
```

Check node status
```bash
./stchaincli status
```

### Your first tx - `send` transaction

```bash
./stchaincli tx send <from address> <to address> <amount> --keyring-backend=<keyring's backend> --chain-id=<current chain-id>
```

```bash
$ ./stchaincli tx send st1qzx8na3ujlaxstgcyguudaecr6mpsemflhhzua st1jvf660xagmzuzyqyqu3w27sj0ragn7qetnwmyr 100000000000ustos --keyring-backend=test --chain-id=stratos-testnet-3 --gas=auto

# then input y for the pop up to confirm send
```
> * In testing phase, --keyring-backend="test"
> * In testing phase, the current `chain-id` may change when updating. When it is applied, user needs to point out `current chain-id` which can be found on [this page](https://big-dipper-test.thestratos.org/), right next to the search bar at the top of the page.

### Becoming a validator(optional)
After the following steps have been done, Any participant in the network can signal that they want to become a validator. Please refer to [How to Become a Validator](https://stratos.gitbook.io/st-docs/stratos-chain-english/stratos-chain-commands/how-to-become-a-validator) for more details about validator creation, delegation as well as FAQ.
- [x] download related files
- [x] start your node to catch up to the latest block height(synchronization)
- [x] create your Stratos Chain Wallet
- [x] `Faucet` or `send` an amount of tokens to this wallet

## Run node as a service
> NOTE: All below steps require *root* privileges

### Create the `/lib/systemd/system/stratos.service` file with the following content

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
> Note:
> In the [service] section, 
> * `User` is your system login user account
> * `Group` is user group
> * `ExecStart` designates the path to the binary file stchaind
> * `--home` designates the path to the node folder.

Please modify these values according to your situations. Make sure the above values are correct.  

### Start service
```
systemctl daemon-reload
systemctl enable stratos.service
systemctl start stratos.service
```

### Check if service is running as expected
```
systemctl status stratos.service
```

### Check the node log
```
journalctl -u stratos.service -f 
# exit with ctrl+c
```
## Documents
More details are available at [here](https://stratos.gitbook.io/st-docs/stratos-chain-english/stratos-chain-testnet/stratos-chain-testnet)

For Chinese version, please refer to [Stratos-chain testnet 测试网说明](https://stratos.gitbook.io/st-docs/stratoschain-zhong-wen-ban/stratoschain-ce-shi-wang/stratoschain-testnet-ce-shi-wang-shuo-ming)
