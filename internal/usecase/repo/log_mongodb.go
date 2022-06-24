package repo

import (
	"mlinaa/internal/entity"
	"mlinaa/pkg/mongodb"
)

type LogRepo struct {
	*mongodb.Mongodb
}

func (l LogRepo) Store(log entity.Log) (status bool, err error) {
	c := l.Builder.DB("mlinna").C("logs")
	err = c.Insert(log)
	if err != nil {
		return true, nil
	}
	return false, err
}

func (l LogRepo) Get(UniqueId string) (log entity.Log) {
	c := l.Builder.DB("mlinna").C("logs")
	
}

func (l LogRepo) Delete(UniqueId string) (status bool, err error) {
	//TODO implement me
	panic("implement me")
}

func New(db *mongodb.Mongodb) *LogRepo {
	return &LogRepo{db}
}
