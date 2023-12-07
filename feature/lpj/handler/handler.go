package handler

import (
	"dami-api/feature/lpj"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service lpj.ServiceLpjInterface
	rg      *gin.RouterGroup
}

func (h *handler) get(ctx *gin.Context) {
	lpj, err := h.service.Select()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res := LpjResponseMulti(lpj)

	ctx.JSON(http.StatusOK, res)
}

func (h *handler) create(ctx *gin.Context) {
	var payloadLpj RequestLpj

	if err := ctx.ShouldBind(&payloadLpj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "harap masukan data dengan benar. Ket :" + err.Error(),
		})
		return
	}

	lpj := LpjRequset(payloadLpj)

	res, err := h.service.Create(lpj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "harap masukan data dengan benar. Ket :" + err.Error(),
		})
		return
	}

	lpjRes := LpjResponse(res)

	ctx.JSON(http.StatusCreated, lpjRes)
}

func (h *handler) update(ctx *gin.Context) {
	var payloadLpj RequestLpj

	if err := ctx.ShouldBind(&payloadLpj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "harap masukan data dengan benar",
		})
		return
	}

	lpj := LpjRequset(payloadLpj)
	lpj.Id = ctx.Param("id")

	res, err := h.service.Update(lpj)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "harap masukan data dengan benar " + err.Error(),
		})
		return
	}

	lpjRes := LpjResponse(res)

	ctx.JSON(http.StatusOK, lpjRes)
}

func (h *handler) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "masukan id sebagai parameter",
		})
		return
	}

	res, err := h.service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "harap masukan data dengan benar " + err.Error(),
		})
		return
	}

	lpjRes := LpjResponse(res)
	ctx.JSON(http.StatusOK, lpjRes)
}

func (h *handler) Router() {
	v2 := h.rg.Group("lpj")

	v2.POST("", h.create)
	v2.GET("", h.get)
	v2.PUT("/:id", h.update)
	v2.DELETE("/:id", h.delete)
}

func NewHandler(service lpj.ServiceLpjInterface, rg *gin.RouterGroup) *handler {
	return &handler{
		service: service,
		rg:      rg,
	}
}
