package app

import (
	"dami-api/app/database"
	"dami-api/feature/lpj"
	"dami-api/feature/lpj/data"
	lpjHand "dami-api/feature/lpj/handler"
	lpjserv "dami-api/feature/lpj/service"
	"dami-api/feature/tpb"
	tbp "dami-api/feature/tpb/data"
	tbpHand "dami-api/feature/tpb/handler"
	"dami-api/feature/tpb/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type serverRequirement struct {
	engine *gin.Engine
	db     *sql.DB
	lpj    lpj.ServiceLpjInterface
	tbp    tpb.ServiceTbpInterface
}

func (sr *serverRequirement) SetUpRouter() {
	rg := sr.engine.Group("api/v1")

	tbpHand.NewHandler(sr.tbp, rg).Router()
	lpjHand.NewHandler(sr.lpj, rg).Router()
}

func (sr *serverRequirement) Run() {
	sr.SetUpRouter()

	if err := sr.engine.Run(":8081"); err != nil {
		panic(err)
	}
}

func Server() *serverRequirement {
	eng := gin.Default()
	db := database.ConnectDb()

	lpjData := data.NewLpjQuery(db)
	tbpdata := tbp.NewTbp(db)

	servTbp := service.NewTbpService(tbpdata)
	servLpj := lpjserv.NewLpjService(lpjData, tbpdata)

	return &serverRequirement{
		lpj:    servLpj,
		tbp:    servTbp,
		engine: eng,
		db:     db,
	}
}
