package repository

import (
	"assigment2/service/module/request_order/dto"
	"context"
)

func (r RequestOrderRepository) CreateRequetOrder(ctx context.Context, req dto.CreateRequestOrderRequest) (lastInsertID int) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Raw(
		``,
	).Scan(&lastInsertID)

	if result.Error != nil {
		return
	}

	return
}
