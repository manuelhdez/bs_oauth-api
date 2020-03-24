package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/manuelhdez/bs_oauth-api/src/domains/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHander struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHander{
		service: service,
	}
}

func (handler *accessTokenHander) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusNotImplemented, gin.H{
		"token": accessToken,
	})
}
