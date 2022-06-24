package usecase

import (
	"mlinaa/internal/entity"
)

type (
	LogInterface interface {
		Set(request entity.StoreRequestBody) (UniqueId string)
		GetById(logId string) (log entity.Log)
		List()
		Delete(logId string) (status bool)
	}

	LogRepo interface {
		Store(log entity.Log) (status bool, err error)
		Get(uniqueId string) (log entity.Log)
		Destroy(uniqueId string) (status bool)
	}
)
