package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"webServer/packages/settings"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * settings.PageSize
	}

	return result
}
