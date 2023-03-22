package main

import (
	"GIN_rest_api/controller/productcontroller"
	"GIN_rest_api/models"

	"github.com/gin-gonic/gin"
)

func main(){

r := gin.Default()
models.ConnectDatabase()
r.GET("/api/Products", productcontroller.Index)
r.GET("/api/Products/:Id", productcontroller.Show)
r.POST("/api/Products", productcontroller.Create)
r.PUT("/api/Products/:Id", productcontroller.Update)
r.DELETE("/api/Products", productcontroller.Delete)

r.Run()

}