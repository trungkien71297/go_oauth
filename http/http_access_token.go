package http

import (
	"github.com/gin-gonic/gin"
	"github.com/trungkien71297/go_oauth/domain/access_token"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
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

func (ac *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		c.String(123, "asd")
		return
	}

	if err := ac.service.Create(at); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(200, at)
}
func (ac *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
}
