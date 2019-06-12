package key

import "github.com/wemeetagain/go-hdwallet"

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
func NewMasterKey(seed []byte) MasterKey {
	master := hdwallet.MasterKey(seed)
	return &bip32MasterKey{master}
}
