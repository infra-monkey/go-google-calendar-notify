package gcal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// (GET /)
func (Server) GetHome(ctx *gin.Context) {
	resp := "Welcome"

	ctx.JSON(http.StatusOK, resp)
}
