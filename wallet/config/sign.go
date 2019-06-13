package config

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

// Sign ...
func Sign(message string, priv *rsa.PrivateKey) (string, error) {
	h := crypto.SHA256.New()
	h.Write([]byte(message))
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sign), err
}

// Verify ...
func Verify(message, sign string, pub *rsa.PublicKey) error {
	h := crypto.SHA256.New()
	h.Write([]byte(message))
	hashed := h.Sum(nil)

	dst := make([]byte, base64.StdEncoding.DecodedLen(len(sign)))
	n, err := base64.StdEncoding.Decode(dst, []byte(sign))
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, dst[:n])
}
