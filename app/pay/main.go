package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"zf/controllers"
	"zf/service"
)

func init() {
	var err error
	service.Db, err = service.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
}

func main() {
	r := gin.Default()

	r.POST("/payments", controllers.CreatePayment)
	r.PUT("/payments/:payment_id/cancel", controllers.CancelPayment)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
