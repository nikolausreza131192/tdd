package main_test

import (
	"errors"
	"testing"

	. "github.com/nikolausreza131192/tdd/advanced"
	"github.com/stretchr/testify/assert"
)

func TestBuy(t *testing.T) {
	tcs := []struct {
		name              string
		cart              Cart
		payment           Payment
		mockOrderDB       func(order Order) (int, error)
		mockOrderDetailDB func(orderDetail OrderDetail) error
		mockPaymentDB     func(payment Payment) error
		expectedError     error
	}{
		// If cart ID is empty, then should return error
		{
			name: "Cart ID is empty; Should return error",
			cart: Cart{
				Qty:    10,
				ItemID: 1,
				Price:  5000,
				UserID: 8966442,
			},
			expectedError: errors.New("Invalid cart ID"),
		},
		// If qty is 0, then should return error
		{
			name: "Qty is 0; Should return error;",
			cart: Cart{
				CartID: 1,
				Qty:    0,
				ItemID: 1,
				Price:  5000,
				UserID: 8966442,
			},
			expectedError: errors.New("Invalid qty"),
		},
		// If item ID is empty, then should return error
		{
			name: "ItemID is 0; Should return error;",
			cart: Cart{
				CartID: 1,
				Qty:    10,
				Price:  5000,
				UserID: 8966442,
			},
			expectedError: errors.New("Invalid itemID"),
		},
		// Payment amount is not equal with cart amount, then should return error
		{
			name: "Payment amount is not equal; Should return error",
			cart: Cart{
				CartID: 1,
				Qty:    10,
				ItemID: 3,
				Price:  5000,
				UserID: 8966442,
			},
			payment: Payment{
				Amount: 1000,
			},
			expectedError: errors.New("Invalid payment amount"),
		},
		// Store order to DB error, then should return error
		{
			name: "Store order to DB error;",
			cart: Cart{
				CartID: 1,
				Qty:    10,
				ItemID: 3,
				Price:  5000,
				UserID: 8966442,
			},
			payment: Payment{
				Amount: 50000,
			},
			mockOrderDB: func(order Order) (int, error) {
				return 0, errors.New("Failed to save order to DB")
			},
			expectedError: errors.New("Failed to save order to DB"),
		},
		// Store order detail to DB error, then should return error
		{
			name: "Store order detail to DB error;",
			cart: Cart{
				CartID: 1,
				Qty:    10,
				ItemID: 3,
				Price:  5000,
				UserID: 8966442,
			},
			payment: Payment{
				Amount: 50000,
			},
			mockOrderDB: func(order Order) (int, error) {
				return 99, nil
			},
			mockOrderDetailDB: func(orderDetail OrderDetail) error {
				return errors.New("Failed to save order detail to DB")
			},
			expectedError: errors.New("Failed to save order detail to DB"),
		},
		// Store payment data to DB error, then should return error
		{
			name: "Store payment data to DB error;",
			cart: Cart{
				CartID: 1,
				Qty:    10,
				ItemID: 3,
				Price:  5000,
				UserID: 8966442,
			},
			payment: Payment{
				Amount: 50000,
			},
			mockOrderDB: func(order Order) (int, error) {
				return 99, nil
			},
			mockOrderDetailDB: func(orderDetail OrderDetail) error {
				return nil
			},
			mockPaymentDB: func(payment Payment) error {
				return errors.New("Failed to save payment data to DB")
			},
			expectedError: errors.New("Failed to save payment data to DB"),
		},
		// Success
		{
			name: "Success",
			cart: Cart{
				CartID: 1,
				Qty:    10,
				ItemID: 3,
				Price:  5000,
				UserID: 8966442,
			},
			payment: Payment{
				Amount: 50000,
			},
			mockOrderDB: func(order Order) (int, error) {
				return 99, nil
			},
			mockOrderDetailDB: func(orderDetail OrderDetail) error {
				return nil
			},
			mockPaymentDB: func(payment Payment) error {
				return nil
			},
			expectedError: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			SaveOrderToDB = tc.mockOrderDB
			SaveOrderDetailToDB = tc.mockOrderDetailDB
			SavePaymentToDB = tc.mockPaymentDB
			err := Buy(tc.cart, tc.payment)

			assert.Equal(t, tc.expectedError, err)
		})
	}
}
