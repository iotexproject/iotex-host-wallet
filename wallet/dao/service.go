package dao

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"

	conf "github.com/iotexproject/iotex-host-wallet/wallet/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service service model
type Service struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	Name      string        `bson:"name" json:"name"`
	APIKey    string        `bson:"api_key" json:"api_key"`
	Status    string        `bson:"status" json:"status"`
	PublicKey string        `bson:"public_key" json:"public_key"`
	CreatedAt int64         `bson:"created_at" json:"created_at"`
	UpdatedAt int64         `bson:"updated_at" json:"updated_at"`
}

// ServiceFindByAPIKey find service by apikey
func ServiceFindByAPIKey(apikey string) (*Service, error) {
	var service Service
	c := conf.C.Container.MongoSession.DB("").C("service")
	err := c.Find(bson.M{"api_key": apikey}).One(&service)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &service, nil
}

// GetPublicKey get rsa PublicKey
func (s *Service) GetPublicKey() (*rsa.PublicKey, error) {
	dst, err := base64.StdEncoding.DecodeString(s.PublicKey)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(dst)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}
