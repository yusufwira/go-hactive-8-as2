package repository

import (
	"assigment2/service/module/request_order/dto"
	"assigment2/service/module/request_order/entity"
	"context"
	"fmt"
	"log"
)

func (r RequestOrderRepository) CreateRequetOrder(ctx context.Context, req entity.RequestOrder) (lastInsertID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	newRo := entity.RequestOrder{
		CustomerName: req.CustomerName,
		OrderAt:      req.OrderAt,
	}
	result := db.Create(&newRo) // Insert into the database
	if result.Error != nil {
		log.Fatal("Error inserting data:", result.Error)
		return
	}

	lastInsertID = newRo.OrderId

	return
}

func (r RequestOrderRepository) CreateItem(ctx context.Context, req entity.Items) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Create(&req) // Insert into the database
	if result.Error != nil {
		log.Fatal("Error inserting data:", result.Error)
		return
	}
	return
}

func (r RequestOrderRepository) GetAllData(ctx context.Context) (res []dto.RequestOrder, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	temps := []entity.RequestOrder{}
	_ = db.Raw(
		`SELECT * FROM request_orders`,
	).Scan(&temps)
	for _, x := range temps {
		r := dto.RequestOrder{}
		items := []entity.Items{}
		_ = db.Raw(`SELECT * FROM items where order_id = ?`, x.OrderId).Scan(&items)
		r.CustomerName = x.CustomerName
		r.Id = int(x.OrderId)
		r.CreatedAt = x.CreatedAt
		r.UpdatedAt = x.UpdatedAt
		r.Items = items
		res = append(res, r)
	}

	return res, err
}

func (r RequestOrderRepository) GetAllDataById(ctx context.Context, id string) (res dto.RequestOrder, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	temps := entity.RequestOrder{}
	_ = db.Raw(
		`SELECT * FROM request_orders WHERE order_id = ?`, id,
	).Scan(&temps)

	items := []entity.Items{}
	_ = db.Raw(`SELECT * FROM items where order_id = ?`, temps.OrderId).Scan(&items)
	res.CustomerName = temps.CustomerName
	res.Id = int(temps.OrderId)
	res.CreatedAt = temps.CreatedAt
	res.UpdatedAt = temps.UpdatedAt
	res.Items = items

	return res, err
}

func (r RequestOrderRepository) UpdateDataRo(ctx context.Context, id string, req dto.CreateRequestOrderRequest) (lastInsertID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Raw(
		`UPDATE public.request_orders SET customer_name = ? WHERE order_id = ?`,
		req.CustomerName, id,
	).Scan(&lastInsertID)
	fmt.Println(req.CustomerName)

	return
}

func (r RequestOrderRepository) DeleteAllItem(ctx context.Context, id string) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Exec(
		`DELETE FROM public.items WHERE order_id = ?`, id,
	)
	return
}

func (r RequestOrderRepository) DeleteRo(ctx context.Context, id string) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Exec(
		`DELETE FROM public.items WHERE order_id = ?`, id,
	)

	_ = db.Exec(
		`DELETE FROM public.request_orders WHERE order_id = ?`, id,
	)
	return
}
