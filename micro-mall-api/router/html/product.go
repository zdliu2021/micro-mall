package html

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login/index.html", gin.H{})
}
