package main

import "fmt"

// type OrderStatus string

// const (
// 	Received  OrderStatus = "received"
// 	Confirmed             = "confirmed"
// 	Prepared              = "prepared"
// 	Delivered             = "delivered"
// )

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func (s OrderStatus) getName() string {
	switch s {
	case Received:
		return "Received"
	case Confirmed:
		return "Confirmed"
	case Prepared:
		return "Prepared"
	case Delivered:
		return "Delivered"
	default:
		return "Unknown"
	}
}

// String method can also be used to get the string representation of the enum
// func (s OrderStatus) String() string {
// 	switch s {
// 	case Received:
// 		return "Received"
// 	case Confirmed:
// 		return "Confirmed"
// 	case Prepared:
// 		return "Prepared"
// 	case Delivered:
// 		return "Delivered"
// 	default:
// 		return "Unknown"
// 	}
// }

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
	fmt.Println("Order status string representation:", status.getName())
}

func main() {
	changeOrderStatus(Received)
	changeOrderStatus(Confirmed)
	changeOrderStatus(Prepared)
	changeOrderStatus(Delivered)
}
