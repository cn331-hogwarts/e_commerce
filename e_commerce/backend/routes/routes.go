package routers

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(inRoutes *gin.Engine) {
	inRoutes.POST("users/signup")
	inRoutes.POST("users/login")
	inRoutes.POST("/admin/addproduct")
	inRoutes.GET("/users/productview")
	inRoutes.GET("/users/search")
}
