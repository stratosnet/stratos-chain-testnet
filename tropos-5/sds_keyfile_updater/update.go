package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/stratosnet/sds_keyfile_updater/handler"
)

func main() {

	var filePath string
	var password string

	flag.StringVar(&filePath, "file", "", "wallet key file path")
	flag.StringVar(&password, "password", "", "password of wallet key")
	flag.Parse()

	fileDir, originalAddrStr := handler.ParseFileDir(filePath)

	// validate keyfile by name, only wallet key file is allowed to be updated
	if !strings.HasPrefix(originalAddrStr, "st") || len(originalAddrStr) > 41 {
		panic("file is not a valid wallet keyfile.")
	}

	//read old wallet account key file
	accountKey, encryptedKeyJSONV3, err := handler.ReadKeyFile(filePath, password)
	if err != nil {
		panic(err)
	}

	//update address
	encryptedKeyJSONV3.Address = hex.EncodeToString(accountKey.Address.Bytes())

	//get bech32 address string for filename
	bech32Addr, err := bech32.ConvertAndEncode("st", accountKey.Address.Bytes())
	if err != nil {
		panic(err)
	}

	if originalAddrStr == bech32Addr {
		fmt.Println("The key file is already the latest version.")
		return
	}

	//rename old key file
	err = os.Rename(filePath, filePath+".old")
	if err != nil {
		panic(err)
	}

	//write new key file
	newFilePath := fileDir + bech32Addr + ".json"
	err = handler.WriteKeyFile(newFilePath, encryptedKeyJSONV3)
	if err != nil {
		panic(err)
	}

	fmt.Println(bech32Addr + ".json created.")
}
