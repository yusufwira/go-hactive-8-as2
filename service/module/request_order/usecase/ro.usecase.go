package usecase

import (
	"assigment2/service/module/request_order/dto"
	"assigment2/service/module/request_order/entity"
	"context"
	"strconv"
)

func (u *RequestOrderUsecase) CreateRequestOrder(ctx context.Context, req dto.CreateRequestOrderRequest) (res uint, err error) {
	u.GormDB.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		if err != nil {
			u.GormDB.RollbackTransaction()
			return
		}
		u.GormDB.CommitTransaction()
	}()

	requestOrderEntity := entity.RequestOrder{}
	requestOrderEntity.CustomerName = req.CustomerName
	requestOrderEntity.OrderAt = &req.OrderedAt

	idRo, err := u.RequestOrderRepository.CreateRequetOrder(ctx, requestOrderEntity)
	if err != nil {
		return
	} else {
		for i := 0; i < len(req.Items); i++ {
			itemEntity := entity.Items{}
			itemEntity.ItemCode = req.Items[i].ItemCode
			itemEntity.Description = req.Items[i].Description
			itemEntity.Quantity = req.Items[i].Quantity
			itemEntity.OrderId = idRo
			err1 := u.RequestOrderRepository.CreateItem(ctx, itemEntity)
			if err1 != nil {
				return
			}
		}
	}
	u.cache.Delete("requestOrder")

	return idRo, err

}

func (u *RequestOrderUsecase) GetAllData(ctx context.Context) (res []dto.RequestOrder, err error) {

	if cachedRo, found := u.cache.Load("requestOrder"); found {
		return cachedRo.([]dto.RequestOrder), nil
	}
	res, err = u.RequestOrderRepository.GetAllData(ctx)
	if err != nil {
		return
	}

	u.cache.Store("requestOrder", res)

	return res, err
}

func (u *RequestOrderUsecase) UpdateData(ctx context.Context, id string, req dto.CreateRequestOrderRequest) (res dto.RequestOrder, err error) {
	_, err = u.RequestOrderRepository.UpdateDataRo(ctx, id, req)
	if err != nil {
		return
	}

	u.RequestOrderRepository.DeleteAllItem(ctx, id)
	newid, _ := strconv.Atoi(id)
	for i := 0; i < len(req.Items); i++ {
		itemEntity := entity.Items{}
		itemEntity.ItemCode = req.Items[i].ItemCode
		itemEntity.Description = req.Items[i].Description
		itemEntity.Quantity = req.Items[i].Quantity
		itemEntity.OrderId = uint(newid)
		err1 := u.RequestOrderRepository.CreateItem(ctx, itemEntity)
		if err1 != nil {
			return
		}
	}

	data, err := u.RequestOrderRepository.GetAllDataById(ctx, id)

	return data, err
}

func (u *RequestOrderUsecase) DeleteData(ctx context.Context, id string) (err error) {
	err = u.RequestOrderRepository.DeleteRo(ctx, id)
	if err != nil {
		return
	}
	return
}
