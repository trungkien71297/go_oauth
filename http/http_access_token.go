package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/trungkien71297/go_oauth/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (ac *accessTokenHandler) GetById(c *gin.Context) {

	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := ac.service.GetById(accessTokenId)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}
