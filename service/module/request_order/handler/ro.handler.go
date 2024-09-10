package handler

import (
	"assigment2/service/module/request_order/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h RequestOrderHandler) Create(ctx *gin.Context) {
	req := dto.CreateRequestOrderRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.RequestOrderUsecase.CreateRequestOrder(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
			"data":    err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    res,
	})

}

func (h RequestOrderHandler) GetAllData(ctx *gin.Context) {
	res, err := h.RequestOrderUsecase.GetAllData(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
			"data":    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    res,
	})
}

func (h RequestOrderHandler) Update(ctx *gin.Context) {
	ids := ctx.Param("id")
	req := dto.CreateRequestOrderRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := h.RequestOrderUsecase.UpdateData(ctx, ids, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
			"data":    err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    res,
	})
}

func (h RequestOrderHandler) Delete(ctx *gin.Context) {
	ids := ctx.Param("id")

	err := h.RequestOrderUsecase.DeleteData(ctx, ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}
