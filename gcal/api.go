package gcal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /api/status)
func (Server) GetApiStatus(ctx *gin.Context) {
	resp := Status{
		Status: "OK",
	}

	ctx.JSON(http.StatusOK, resp)
}

// (GET /api/accounts)
func (Server) GetApiAccounts(ctx *gin.Context) {
	resp := Accounts{}

	ctx.JSON(http.StatusOK, resp)
}

// (POST /api/accounts)
func (Server) PostApiAccounts(ctx *gin.Context) {
	resp := Accounts{}

	ctx.JSON(http.StatusOK, resp)
}
