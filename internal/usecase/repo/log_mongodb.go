package repo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"mlinaa/internal/entity"
	"mlinaa/pkg/mongodb"
)

type LogRepo struct {
	*mongodb.Mongodb
}

const tableName = "logs"

func New(db *mongodb.Mongodb) *LogRepo {
	return &LogRepo{db}
}

func (l LogRepo) Store(log entity.Log) (status bool, err error) {
	c := l.Builder.C(tableName)
	err = c.Insert(log)
	if err != nil {
		return true, nil
	}
	return false, err
}

func (l LogRepo) Get(uniqueId string) (log entity.Log) {
	result := entity.Log{}

	err := l.Builder.C(tableName).
		Find(bson.M{"unique_id": uniqueId}).One(&result)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return result

}

func (l LogRepo) Destroy(uniqueId string) (status bool) {

	err := l.Builder.C(tableName).Remove(bson.M{"unique_id": uniqueId})
	if err != nil {
		return false
	}
	return true
}
