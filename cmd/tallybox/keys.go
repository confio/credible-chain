package main

import (
	"fmt"
	"os"
	"path/filepath"

	wc "github.com/confio/credible-chain/weaveclient"
)

// KeysCmd loads or creates the private key file.
// It then prints out the address on stdout
// It is safe to run many times, only the first run writes data
func KeysCmd(home string) error {
	err := os.MkdirAll(home, 0700)
	if err != nil {
		return err
	}
	filename := filepath.Join(home, "notary.pk")
	privkey, err := wc.LoadPrivateKey(filename)
	if err != nil {
		privkey = wc.GenPrivateKey()
		err = wc.SavePrivateKey(privkey, filename, false)
		if err != nil {
			return err
		}
	}
	fmt.Printf("%s", privkey.PublicKey().Address())
	return nil
}
