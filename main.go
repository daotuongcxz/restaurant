package main

import (
	"go_project/common"
	appctx "go_project/component"
	"go_project/module/restaurant/transport/ginrestaurant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := common.ConnectDatabase()
	if err != nil {
		return
	}

	// http server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//==================================================//
	// API
	appContext := appctx.NewAppContext(db)
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	restaurants.GET("/:id", ginrestaurant.FindRestaurant(appContext))
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	r.Run() // listen and serve on 0.0.0.0:8080

}
