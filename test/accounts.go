package test

import (
	"bytes"
	"crypto/ecdsa"
	"embed"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	//go:embed testdata/root.key
	rootKeyBytes []byte
	rootKey      = check(crypto.HexToECDSA(strings.TrimSpace(string(rootKeyBytes))))
	rootAddress  = crypto.PubkeyToAddress(rootKey.PublicKey)

	//go:embed testdata/account-?.key
	accountKeys embed.FS

	AccountKeyFile []string
	AccountKey     []*ecdsa.PrivateKey
	AccountAddress []common.Address
)

func init() {
	cwd := check(os.Getwd())
	dir := packageDir()
	entries := check(accountKeys.ReadDir("testdata"))
	for _, entry := range entries {
		keyFile := check(filepath.Rel(cwd, filepath.Join(dir, "testdata", entry.Name())))
		keyData := bytes.TrimSpace(check(accountKeys.ReadFile(path.Join("testdata", entry.Name()))))
		key := check(crypto.HexToECDSA(string(keyData)))
		AccountKeyFile = append(AccountKeyFile, keyFile)
		AccountKey = append(AccountKey, key)
		AccountAddress = append(AccountAddress, crypto.PubkeyToAddress(key.PublicKey))
	}
}
