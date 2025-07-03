package main

import (
	"fmt"
	"time"
)

// we can do object oriented programming in go with the help of struct

type customer struct{
	name string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time    // nanosecond precision
	customer               // struct embedding
}

// constructor-type function
func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type
func (o *order)changeStatus(status string) {
    o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	language:= struct{      // inline struct
		name string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// if we don't set any field then default value is zeroed value i.e int->0, float->0, string->"", bool->false
    myOrders:= newOrder("1",45,"received")
	fmt.Println(myOrders.amount)

	// we can create new struct of customer as well 
	newCustomer:= customer{
		name: "Shubham",
		phone: "1234567890",
	}
    fmt.Println(newCustomer)

    myOrder:= order{
		id: "1",
		amount: 50.00,
		status: "pending",
		customer: customer{                // in-line
			name: "Shubham",
			phone: "1234567890",
		},
	}
    fmt.Println(myOrder.customer)
	myOrder.customer.name = "Shubham Verma"
	fmt.Println(myOrder.getAmount())

	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder.amount)
	
	fmt.Println("Order Struct: ", myOrder)
}