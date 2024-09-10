package repository

import (
	"assigment2/connection"
	"assigment2/service/module/request_order/dto"
	"assigment2/service/module/request_order/entity"
	"context"
)

type RequestOrderRepository struct {
	gormDB *connection.GormDB
}

func InitRequestOrderRepository(gormDB *connection.GormDB) IRequestOrderRepository {
	return &RequestOrderRepository{
		gormDB: gormDB,
	}
}

type IRequestOrderRepository interface {
	CreateRequetOrder(ctx context.Context, req entity.RequestOrder) (lastInsertID uint, err error)
	CreateItem(ctx context.Context, req entity.Items) (err error)
	GetAllData(ctx context.Context) (res []dto.RequestOrder, err error)
	GetAllDataById(ctx context.Context, id string) (res dto.RequestOrder, err error)
	UpdateDataRo(ctx context.Context, id string, req dto.CreateRequestOrderRequest) (lastInsertID uint, err error)
	DeleteAllItem(ctx context.Context, id string) (err error)
	DeleteRo(ctx context.Context, id string) (err error)
}
