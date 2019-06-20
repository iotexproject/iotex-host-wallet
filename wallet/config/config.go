package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
)

// C config instance
var C Config

// Config ...
type Config struct {
	Port  int32  `yaml:"port"`
	Mongo string `yaml:"mongo"`
	Keys  struct {
		Wallet struct {
			PrivateKey string `yaml:"privateKey"`
		} `yaml:"wallet"`
	} `yaml:"keys"`
	Container struct {
		MongoSession     *mgo.Session
		WalletPrivateKey *rsa.PrivateKey
		ServicePublicKey *rsa.PublicKey
	} `yaml:"-"`
}

func init() {
	env := os.Getenv("env")
	f := "config.yaml"
	if env != "" {
		f = "config-" + env + ".yaml"
	}
	fb, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf("read config file %s fail, %v", f, err)
	}

	err = yaml.Unmarshal(fb, &C)
	if err != nil {
		log.Fatalf("parse config file %s fail, %v", f, err)
	}

	sess, err := mgo.Dial(C.Mongo)
	if err != nil {
		log.Panicf("connect mongo %s error: %v", C.Mongo, err)
	}
	C.Container.MongoSession = sess

	priv, err := restorePrivateKey(C.Keys.Wallet.PrivateKey)
	if err != nil {
		log.Fatalf("restore private key error: %v", err)
	}
	C.Container.WalletPrivateKey = priv
}

func restorePrivateKey(pem string) (*rsa.PrivateKey, error) {
	dst, err := base64.StdEncoding.DecodeString(pem)
	if err != nil {
		return nil, err
	}
	key, err := x509.ParsePKCS8PrivateKey(dst)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}

func restorePublicKey(pem string) (*rsa.PublicKey, error) {
	dst, err := base64.StdEncoding.DecodeString(pem)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(dst)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}
