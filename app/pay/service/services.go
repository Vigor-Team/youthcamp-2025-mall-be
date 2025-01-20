package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"zf/models"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// InitDB initializes the database connection.
func InitDB() (*sql.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	fmt.Println("Connected to database successfully.")
	Db = db
	return db, nil
}

// CreatePayment inserts a new payment into the database.
func CreatePayment(payment models.Payment) error {
	query := `
	INSERT INTO payments (id, order_id, payment_method, total_amount, description, status, payment_gateway, payment_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := Db.Exec(query, payment.ID, payment.OrderID, payment.PaymentMethod, payment.TotalAmount, payment.Description, payment.Status, payment.PaymentGateway, payment.PaymentID)
	if err != nil {
		return fmt.Errorf("error creating payment: %w", err)
	}

	return nil
}

// UpdatePaymentStatus updates the status of a payment in the database.
func UpdatePaymentStatus(paymentID string, status string) error {
	query := `
	UPDATE payments
	SET status = ?, updated_at = NOW()
	WHERE payment_id = ? AND status = 'processing'`

	res, err := Db.Exec(query, status, paymentID)
	if err != nil {
		return fmt.Errorf("error updating payment status: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("payment not found or already processed")
	}

	return nil
}

// ScheduleCancelPayment schedules a payment cancellation after a specified duration.
func ScheduleCancelPayment(paymentID string, duration time.Duration) {
	go func() {
		timer := time.NewTimer(duration)
		defer timer.Stop()

		select {
		case <-timer.C:
			err := UpdatePaymentStatus(paymentID, "cancelled")
			if err != nil {
				log.Printf("Failed to cancel payment %s: %v", paymentID, err)
			} else {
				log.Printf("Payment %s cancelled after timeout", paymentID)
			}
		}
	}()
}
