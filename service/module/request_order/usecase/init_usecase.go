package usecase

import (
	"assigment2/connection"
	"assigment2/service/module/request_order/dto"
	roRepo "assigment2/service/module/request_order/repository"
	"context"
	"sync"
)

type RequestOrderUsecase struct {
	RequestOrderRepository roRepo.IRequestOrderRepository
	GormDB                 *connection.GormDB
	cache                  sync.Map
}

func InitRequestOrderUsecase(requestOrderRepository roRepo.IRequestOrderRepository, gormDB *connection.GormDB) IRequestOrderUsecase {
	return &RequestOrderUsecase{
		RequestOrderRepository: requestOrderRepository,
		GormDB:                 gormDB,
	}
}

type IRequestOrderUsecase interface {
	CreateRequestOrder(ctx context.Context, req dto.CreateRequestOrderRequest) (res uint, err error)
	GetAllData(ctx context.Context) (res []dto.RequestOrder, err error)
	UpdateData(ctx context.Context, id string, req dto.CreateRequestOrderRequest) (res dto.RequestOrder, err error)
	DeleteData(ctx context.Context, id string) (err error)
}
