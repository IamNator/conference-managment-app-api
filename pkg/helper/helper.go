package helper

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetBearerToken(ctx *gin.Context) string {
	return strings.TrimSpace(strings.Replace(ctx.GetHeader("Authorization"), "Bearer", "", -1))
}
