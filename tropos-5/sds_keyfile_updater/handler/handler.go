package handler

import (
	"encoding/json"
	"io/ioutil"
)

func WriteKeyFile(filePath string, encryptedKeyJSONV3 *EncryptedKeyJSONV3) error {
	fileBytes, err := json.Marshal(encryptedKeyJSONV3)
	if err != nil {
		return err
	}
	err = WriteFile(filePath, fileBytes)
	if err != nil {
		return err
	}
	return nil
}

func ReadKeyFile(filePath, password string) (*AccountKey, *EncryptedKeyJSONV3, error) {
	keyjson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}
	key, encryptedKeyJSONV3, err := DecryptKey(keyjson, password)

	if err != nil {
		return nil, nil, err
	}
	return key, encryptedKeyJSONV3, nil
}
