package http

import (
	"net/http"
	"strings"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	atDomain "github.com/manuelhdez/bs_oauth-api/src/domains/access_token"
	"github.com/manuelhdez/bs_oauth-api/src/services/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

// func NewHandler(service access_token.Service) AccessTokenHandler {
// 	return &accessTokenHander{
// 		service: service,
// 	}
// }

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": accessToken,
	})
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	accessToken, err := handler.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}
