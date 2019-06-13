package dao

import (
	"errors"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	conf "github.com/iotexproject/iotex-host-wallet/wallet/config"
)

var mux sync.Mutex

// Config index config
type Config struct {
	ID        bson.ObjectId `bson:"_id" json:"-"`
	LastIndex uint32        `bson:"last_index" json:"last_index"`
	CreatedAt int64         `bson:"created_at" json:"created_at"`
	UpdatedAt int64         `bson:"updated_at" json:"updated_at"`
}

// Save config
func (config *Config) Save() error {
	c := conf.C.Container.MongoSession.DB("").C("config")
	if config.ID == "" {
		config.ID = bson.NewObjectId()
		config.CreatedAt = time.Now().Unix()
		return c.Insert(config)
	}
	return c.UpdateId(config.ID, config)
}

// ConfigNewIndex get new index
func ConfigNewIndex() (uint32, error) {
	mux.Lock()
	defer mux.Unlock()
	var config Config
	c := conf.C.Container.MongoSession.DB("").C("config")
	err := c.Find(nil).One(&config)
	if err != nil {
		if err == mgo.ErrNotFound {
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
