package usecase

import (
	"fmt"
	"github.com/rs/xid"
	"mlinaa/internal/entity"
	"strconv"
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

func (l LogUseCase) Set(request entity.StoreRequestBody) string {

	uniqueId := xid.New().String() + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	log := entity.Log{
		Data:        request.Data,
		UniqueId:    uniqueId,
		RequestTime: time.Now(),
	}

	l.repo.Store(log)

	return uniqueId
}

func (l LogUseCase) GetById(logId string) (log entity.Log) {
	fmt.Println(logId)
	return l.repo.Get(logId)
}

func (l LogUseCase) Delete(logId string) (status bool) {

	if l.repo.Destroy(logId) {
		return true
	}

	return false
}

func (l LogUseCase) List() {
	//TODO implement me
	panic("implement me")
}
