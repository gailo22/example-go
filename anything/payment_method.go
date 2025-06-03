package main

import (
	"fmt"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) bool
}

type CreditCard struct{}

func (c CreditCard) ProcessPayment(amount float64) bool {
	fmt.Println("Processing credit card payment...")
	return true
}

type Paypal struct{}

func (p Paypal) ProcessPayment(amount float64) bool {
	fmt.Println("Processing paypal payment...")
	return true
}

func main() {
	creditCard := CreditCard{}
	paypal := Paypal{}

	paymentMethods := []PaymentMethod{creditCard, paypal}

	amount := 100.0
	for _, p := range paymentMethods {
		p.ProcessPayment(amount)
	}
}
