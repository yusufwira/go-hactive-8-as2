package main

import (
	"assigment2/connection"
	roHandler "assigment2/service/module/request_order/handler"
	roRepo "assigment2/service/module/request_order/repository"
	roUsecase "assigment2/service/module/request_order/usecase"

	"github.com/gin-gonic/gin"
)

type requestOrderRoutes struct {
	Handler roHandler.RequestOrderHandler
	Router  *gin.RouterGroup
}

func InitRequestOrderRoute(
	router *gin.RouterGroup, gormDB *connection.GormDB,
) *requestOrderRoutes {
	handler := roHandler.InitRequestOrderHandler(
		roUsecase.InitRequestOrderUsecase(
			roRepo.InitRequestOrderRepository(
				gormDB,
			),
			gormDB,
		),
	)
	return &requestOrderRoutes{
		Handler: *handler,
		Router:  router,
	}
}

func (r *requestOrderRoutes) Routes() {

}
