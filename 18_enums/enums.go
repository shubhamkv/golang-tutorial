package main

import "fmt"

// enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Prepared OrderStatus = "prepared"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)
}