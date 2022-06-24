package mongodb

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
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
	Builder *mgo.Session
}

func New() (*Mongodb, error) {
	//e := godotenv.Load() //Load .env file
	//if e != nil {
	//	panic(e)
	//}

	mongo := &Mongodb{}
	Host := []string{
		"127.0.0.1:27017",
		// replica set addrs...
	}

	info, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DB_NAME"),
		//Username: username,
		//Password: password,
	})

	if err != nil {
		return nil, fmt.Errorf("mongodb - connAttempts == 0: %w", err)
	}

	mongo.Builder = info

	return mongo, nil
}
