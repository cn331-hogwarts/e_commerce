package main

import (
	"github/Augustus2011/ecommerce/controllers"
	"github/Augustus2011/ecommerce/data"
	"github/Augustus2011/ecommerce/middleware"
	"log"
	"os"

	"github.com/gin-goinc/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(data.ProductData(data.Client, "Products"), data.UserData(data.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addcart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.ButFromCart())
	router.GET("instantbut", app.InstantBuy())

	log.Fatal(router.Run(":") + port)
}
