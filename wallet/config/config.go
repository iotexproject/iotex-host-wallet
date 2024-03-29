package config

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

// C config instance
var C Config

// Config ...
type Config struct {
	Port     int32  `yaml:"port"`
	Mongo    string `yaml:"mongo"`
	Database string `yaml:"database"`
	Keys     struct {
		Wallet struct {
			PrivateKey string `yaml:"privateKey"`
		} `yaml:"wallet"`
		Service struct {
			PublicKey string `yaml:"publicKey"`
		} `yaml:"service"`
	} `yaml:"keys"`
	Container struct {
		MongoClient      *mongo.Client
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

	if C.Database == "" {
		C.Database = "hostwallet"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(C.Mongo))
	if err != nil {
		log.Panicf("create mongo client %s error: %v", C.Mongo, err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Panicf("connect mongo %s error: %v", C.Mongo, err)
	}

	C.Container.MongoClient = client
	go func() {
		for {
			err = C.Container.MongoClient.Ping(context.Background(), readpref.Primary())
			if err != nil {
				log.Printf("ping mongo session fail, reconnect...")
				client, err := mongo.NewClient(options.Client().ApplyURI(C.Mongo))
				if err != nil {
					log.Panicf("create mongo client %s error: %v", C.Mongo, err)
				}
				ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
				err = client.Connect(ctx)
				if err != nil {
					log.Panicf("connect mongo %s error: %v", C.Mongo, err)
				}
				C.Container.MongoClient = client
			}
			time.Sleep(time.Minute * 5)
		}
	}()

	priv, err := restorePrivateKey(C.Keys.Wallet.PrivateKey)
	if err != nil {
		log.Fatalf("restore private key error: %v", err)
	}
	C.Container.WalletPrivateKey = priv
	pub, err := restorePublicKey(C.Keys.Service.PublicKey)
	if err != nil {
		log.Fatalf("restore public key error: %v", err)
	}
	C.Container.ServicePublicKey = pub
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
