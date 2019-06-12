package key

import (
	"os"

	"github.com/ququzone/go-common/crypto"
)

// ReadSeed read master seed from chain file
func ReadSeed(file, password string) ([]byte, error) {
	aes := crypto.NewAES(password, file)

	cf, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer cf.Close()

	cb := make([]byte, 272)
	_, err = cf.Read(cb)
	if err != nil {
		return nil, err
	}

	return aes.Decrypt(cb)
}
