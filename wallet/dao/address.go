package dao

import (
	"context"
	"time"

	conf "github.com/iotexproject/iotex-host-wallet/wallet/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Address user address model
type Address struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	UserID    string             `bson:"user_id" json:"user_id"`
	Index     uint32             `bson:"index" json:"index"`
	Address   string             `bson:"address" json:"address"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
	UpdatedAt int64              `bson:"updated_at,omitempty" json:"updated_at"`
}

// Save config
func (address *Address) Save() error {
	c := conf.C.Container.MongoClient.Database(conf.C.Database).Collection("address")

	address.CreatedAt = time.Now().Unix()
	res, err := c.InsertOne(context.Background(), address)
	if err != nil {
		return err
	}
	address.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

// AddressFindByUserID find address by userID
func AddressFindByUserID(userID string) (*Address, error) {
	var address Address
	c := conf.C.Container.MongoClient.Database(conf.C.Database).Collection("address")
	err := c.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&address)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &address, nil
}

// AddressFindByAddress find address by address
func AddressFindByAddress(address string) (*Address, error) {
	var result Address
	c := conf.C.Container.MongoClient.Database(conf.C.Database).Collection("address")
	err := c.FindOne(context.Background(), bson.M{"address": address}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

// AddressFindByAddressAndUserId find address by address and userId
func AddressFindByAddressAndUserId(userID, address string) (*Address, error) {
	var result Address
	c := conf.C.Container.MongoClient.Database(conf.C.Database).Collection("address")
	err := c.FindOne(context.Background(), bson.M{"address": address, "user_id": userID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
