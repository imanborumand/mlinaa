package mongodb

import (
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

func New() (*mgo.Database, error) {
	e := godotenv.Load() //Load .env file
	if e != nil {
		panic(e)
	}

	Host := []string{
		"localhost:27017",
		// replica set addrs...
	}

	info := &mgo.DialInfo{
		Addrs:    Host,
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DB_NAME"),
		//Username: username,
		//Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	return session.DB(os.Getenv("MONGO_DB_NAME")), err
}
