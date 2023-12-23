package routes

import (
	"github/Augustus2011/ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(inRoutes *gin.Engine) {
	inRoutes.POST("users/signup", controllers.Signup())
	inRoutes.POST("users/login", controllers.Login())
	inRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	inRoutes.GET("/users/productview", controllers.SearchProduct())
	inRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
