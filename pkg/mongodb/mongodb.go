package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"mlinaa/config"
	"time"
)

const (
	MongoDBHosts = "localhost:27017"
	AuthDatabase = "mlinaa"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "mlinaa"
)

type Mongodb struct {
	Builder *mgo.Database
}

func New() (*Mongodb, error) {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	mongo := &Mongodb{}
	Host := []string{
		"127.0.0.1:27017",
		// replica set addrs...
	}

	info, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Timeout:  60 * time.Second,
		Database: cfg.MongoDatabase.Name,
		//Username: username,
		//Password: password,
	})

	if err != nil {
		return nil, fmt.Errorf("mongodb - connAttempts == 0: %w", err)
	}

	mongo.Builder = info.DB(cfg.MongoDatabase.Name)

	return mongo, nil
}
