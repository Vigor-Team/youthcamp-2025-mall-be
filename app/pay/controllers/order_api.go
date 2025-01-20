package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"zf/models"
	"zf/service"
)

// CreatePayment handles the creation of a new payment.
func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.ID = models.NewPayment().ID

	err := service.CreatePayment(payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	service.ScheduleCancelPayment(*payment.PaymentID, 5*time.Minute)

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Payment created",
		"payment_id": payment.ID,
	})
}

// CancelPayment handles the cancellation of an existing payment.
func CancelPayment(c *gin.Context) {
	paymentID := c.Param("payment_id")
	err := service.UpdatePaymentStatus(paymentID, "cancelled")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment cancelled",
	})
}
