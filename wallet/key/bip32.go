package key

import "github.com/wemeetagain/go-hdwallet"

// MK singleton master key
var MK MasterKey

// MasterKey master key interface
type MasterKey interface {
	ChildKey(index uint32) ([]byte, error)
}

type bip32MasterKey struct {
	master *hdwallet.HDWallet
}

func (mk *bip32MasterKey) ChildKey(index uint32) ([]byte, error) {
	child, err := mk.master.Child(index)
	if err != nil {
		return nil, err
	}
	return child.Key, nil
}

// NewMasterKey create master key from seed
func NewMasterKey(seed []byte) {
	master := hdwallet.MasterKey(seed)
	MK = &bip32MasterKey{master}
}
