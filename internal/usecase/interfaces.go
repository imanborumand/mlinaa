package usecase

import (
	"context"
	"mlinaa/internal/entity"
)

type (
	LogInterface interface {
		Set(ctx context.Context) (UniqueId string)
		GetById()
		List()
		Forget()
	}

	LogRepo interface {
		Store(log entity.Log) (status bool, err error)
		Get(UniqueId string) (log entity.Log)
		Delete(UniqueId string) (status bool, err error)
	}
)
