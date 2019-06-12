package config

import (
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
	IPs       []string `yaml:"ips"`
	Mongo     string   `yaml:"mongo"`
	Container struct {
		MongoSession *mgo.Session
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
}
