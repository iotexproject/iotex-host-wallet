package dao

import (
	"context"
	"errors"
	"sync"
	"time"

	conf "github.com/iotexproject/iotex-host-wallet/wallet/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mux sync.Mutex

// Config index config
type Config struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	LastIndex uint32             `bson:"last_index" json:"last_index"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
	UpdatedAt int64              `bson:"updated_at" json:"updated_at"`
}

// Save config
func (config *Config) Save() error {
	c := conf.C.Container.MongoClient.Database(conf.C.Database).Collection("config")
	if config.ID.String() == "" {
		config.ID = primitive.NewObjectID()
		config.CreatedAt = time.Now().Unix()
		_, err := c.InsertOne(context.Background(), config)
		return err
	}
	_, err := c.UpdateOne(context.Background(), bson.M{"_id": config.ID}, bson.D{
		{"$set", config},
	})
	return err
}

// ConfigNewIndex get new index
func ConfigNewIndex() (uint32, error) {
	mux.Lock()
	defer mux.Unlock()
	var config Config
	c := conf.C.Container.MongoClient.Database(conf.C.Database).Collection("config")
	err := c.FindOne(context.Background(), bson.M{}).Decode(&config)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			config = Config{LastIndex: 1}
			err := config.Save()
			if err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, errors.New("query config error")
	}
	config.LastIndex = config.LastIndex + 1
	config.UpdatedAt = time.Now().Unix()
	err = config.Save()
	if err != nil {
		return 0, errors.New("save config error")
	}
	return config.LastIndex, nil
}
