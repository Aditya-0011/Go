package main

import "fmt"

type paymentGateway interface {
	pay(amt float32)
}

type payment struct {
	gateway paymentGateway
}

func (p *payment) makePayment(amt float32) {
	p.gateway.pay(amt)
}

type razorpay struct {
}

func (r razorpay) pay(amt float32) {
	fmt.Println("Payment made using Razorpay:", amt)
}

type stripe struct {
}

func (r stripe) pay(amt float32) {
	fmt.Println("Payment made using Stripe:", amt)
}

func main() {
	stripeGateway := stripe{}
	newPayment := payment{
		gateway: stripeGateway,
	}
	newPayment.makePayment(100.50)
}
