package usecase

import (
	"assigment2/connection"
	roRepo "assigment2/service/module/request_order/repository"
)

type RequestOrderUsecase struct {
	RequestOrderRepository roRepo.IRequestOrderRepository
	GormDB                 *connection.GormDB
}

func InitRequestOrderUsecase(requestOrderRepository roRepo.IRequestOrderRepository, gormDB *connection.GormDB) IRequestOrderUsecase {
	return &RequestOrderUsecase{
		RequestOrderRepository: requestOrderRepository,
		GormDB:                 gormDB,
	}
}

type IRequestOrderUsecase interface {
}
