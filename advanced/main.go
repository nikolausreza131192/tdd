package main

import "errors"

func main() {}

type Cart struct {
	CartID int
	Qty    float64
	ItemID int
	Price  float64
	UserID int
}

type Payment struct {
	Amount        float64
	PaymentMethod string
}

type Order struct {
	UserID int
	Total  float64
}

type OrderDetail struct {
	OrderID int
	ItemID  int
	Price   float64
	Qty     float64
	Total   float64
}

var SaveOrderToDB = func(order Order) (int, error) {
	return 0, nil
}

var SaveOrderDetailToDB = func(orderDetailData OrderDetail) error {
	return nil
}

var SavePaymentToDB = func(paymentData Payment) error {
	return nil
}

func Buy(cart Cart, payment Payment) error {
	if cart.CartID == 0 {
		return errors.New("Invalid cart ID")
	} else if cart.Qty == 0 {
		return errors.New("Invalid qty")
	} else if cart.ItemID == 0 {
		return errors.New("Invalid itemID")
	}

	cartAmount := cart.Price * cart.Qty
	if payment.Amount != cartAmount {
		return errors.New("Invalid payment amount")
	}

	orderData := Order{
		UserID: cart.UserID,
		Total:  cartAmount,
	}
	orderID, err := SaveOrderToDB(orderData)
	if err != nil {
		return err
	}

	orderDetailData := OrderDetail{
		OrderID: orderID,
		ItemID:  cart.ItemID,
		Price:   cart.Price,
		Qty:     cart.Qty,
		Total:   cartAmount,
	}
	if err := SaveOrderDetailToDB(orderDetailData); err != nil {
		return err
	}

	if err := SavePaymentToDB(payment); err != nil {
		return err
	}
	return nil
}
