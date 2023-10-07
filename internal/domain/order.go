package domain

import (
	"errors"
)

type OrderStatus string

var (
	ErrAlreadyUploadedByThisUser    = errors.New("the order number has already been uploaded by this user")
	ErrAlreadyUploadedByAnotherUser = errors.New("the order number has already been uploaded by another user")
	ErrIncorrectOrder               = errors.New("incorrect order id")
	ErrNoData                       = errors.New("no response data")
)

const (
	// заказ загружен в систему, но не попал в обработку.
	Registered OrderStatus = "NEW"
	// вознаграждение за заказ рассчитывается.
	Invalid OrderStatus = "PROCESSING"
	// система расчёта вознаграждений отказала в расчёте.
	Processing OrderStatus = "INVALID"
	// данные по заказу проверены и информация о расчёте успешно получена.
	Processed OrderStatus = "PROCESSED"
)

type Order struct {
	OrderID    string      `json:"number"`
	Status     OrderStatus `json:"status"`
	Bonuses    float32     `json:"accrual"`
	UploadedAt string      `json:"uploaded_at"`
	UserID     int64       `json:"-"`
}