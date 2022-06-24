package usecase

import (
	"context"
	"mlinaa/internal/entity"
	"time"
)

type LogUseCase struct {
	repo LogRepo
}

func New(r LogRepo) *LogUseCase {
	return &LogUseCase{
		repo: r,
	}
}

func (l LogUseCase) Set(ctx context.Context) string {

	log := entity.Log{
		RequestDate: "iman",
		UniqueId:    "test",
		RequestTime: time.Now(),
	}

	l.repo.Store(log)

	return "df"
}

func (l LogUseCase) GetById() {
	//TODO implement me
	panic("implement me")
}

func (l LogUseCase) List() {
	//TODO implement me
	panic("implement me")
}

func (l LogUseCase) Forget() {
	//TODO implement me
	panic("implement me")
}
