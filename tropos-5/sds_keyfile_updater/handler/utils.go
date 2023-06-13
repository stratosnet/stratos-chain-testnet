package handler

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ParseFileDir(fileURI string) (dir, originalAddrStr string) {
	dir, fileName := filepath.Split(fileURI)
	fileSuffix := filepath.Ext(fileName)
	originalAddrStr = strings.TrimSuffix(fileName, fileSuffix)
	return dir, originalAddrStr
}

func WriteFile(file string, content []byte) error {
	// Create the keystore directory with appropriate permissions
	// in case it is not present yet.
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return err
	}
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()
	return os.Rename(f.Name(), file)
}
