package dao

import (
	"time"

	conf "github.com/iotexproject/iotex-host-wallet/wallet/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Address user address model
type Address struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	UserID    string        `bson:"user_id" json:"user_id"`
	Index     uint32        `bson:"index" json:"index"`
	Address   string        `bson:"address" json:"address"`
	CreatedAt int64         `bson:"created_at" json:"created_at"`
	UpdatedAt int64         `bson:"updated_at" json:"updated_at"`
}

// Save config
func (address *Address) Save() error {
	c := conf.C.Container.MongoSession.DB("").C("address")
	address.ID = bson.NewObjectId()
	address.CreatedAt = time.Now().Unix()
	return c.Insert(address)
}

// AddressFindByUserID find address by userID
func AddressFindByUserID(userID string) (*Address, error) {
	var address Address
	c := conf.C.Container.MongoSession.DB("").C("address")
	err := c.Find(bson.M{"user_id": userID}).One(&address)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &address, nil
}
