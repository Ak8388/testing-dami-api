package handler

import (
	"dami-api/feature/tpb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerTbp struct {
	service tpb.ServiceTbpInterface
	rg      *gin.RouterGroup
}

func (h *handlerTbp) get(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := h.service.SelectTbpById(id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *handlerTbp) Router() {
	h.rg.GET("/:id", h.get)
}

func NewHandler(service tpb.ServiceTbpInterface, rg *gin.RouterGroup) *handlerTbp {
	return &handlerTbp{
		service: service,
		rg:      rg,
	}
}
