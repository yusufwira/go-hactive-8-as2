package handler

import (
	roUsecase "assigment2/service/module/request_order/Usecase"
)

type RequestOrderHandler struct {
	RequestOrderUsecase roUsecase.IRequestOrderUsecase
}

func InitRequestOrderHandler(roUsecase roUsecase.IRequestOrderUsecase) *RequestOrderHandler {
	return &RequestOrderHandler{
		RequestOrderUsecase: roUsecase,
	}
}
