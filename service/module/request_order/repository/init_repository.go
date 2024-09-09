package repository

import (
	"assigment2/connection"
	"assigment2/service/module/request_order/dto"
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
	CreateRequetOrder(ctx context.Context, req dto.CreateRequestOrderRequest) (lastInsertID int)
}
