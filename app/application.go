package app

import (
	"github.com/gin-gonic/gin"
	"github.com/trungkien71297/go_oauth/domain/access_token"
	"github.com/trungkien71297/go_oauth/http"
	"github.com/trungkien71297/go_oauth/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.New())
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
