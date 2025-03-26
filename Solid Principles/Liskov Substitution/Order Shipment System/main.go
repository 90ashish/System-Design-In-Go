package main

import (
	"fmt"
)

// Order represents an order with minimal details.
type Order struct {
	ID     string
	Weight float64
}

// Shipping interface defines a method for shipping an order.
type Shipping interface {
	Ship(order Order) string
}

// StandardShipping implements the Shipping interface.
type StandardShipping struct{}

func (s StandardShipping) Ship(order Order) string {
	return fmt.Sprintf("Order %s shipped via Standard Shipping", order.ID)
}

// ExpeditedShipping implements the Shipping interface.
type ExpeditedShipping struct{}

func (e ExpeditedShipping) Ship(order Order) string {
	return fmt.Sprintf("Order %s shipped via Expedited Shipping", order.ID)
}

// InternationalShipping implements the Shipping interface.
type InternationalShipping struct {
	Country string
}

func (i InternationalShipping) Ship(order Order) string {
	return fmt.Sprintf("Order %s shipped internationally to %s", order.ID, i.Country)
}

// OrderProcessor handles order processing and delegates shipping via a Shipping interface.
type OrderProcessor struct {
	shippingMethod Shipping
}

// NewOrderProcessor creates a new OrderProcessor with the provided Shipping method.
func NewOrderProcessor(shippingMethod Shipping) *OrderProcessor {
	return &OrderProcessor{shippingMethod: shippingMethod}
}

// ProcessOrder processes an order by using the injected Shipping method.
func (op *OrderProcessor) ProcessOrder(order Order) {
	fmt.Println("Processing order:", order.ID)
	shippingInfo := op.shippingMethod.Ship(order)
	fmt.Println(shippingInfo)
}

func main() {
	// Example order
	order := Order{ID: "ORD123", Weight: 2.5}

	// Process order using Standard Shipping
	standardShipping := StandardShipping{}
	processor := NewOrderProcessor(standardShipping)
	processor.ProcessOrder(order)

	// Process order using Expedited Shipping
	expeditedShipping := ExpeditedShipping{}
	processor = NewOrderProcessor(expeditedShipping)
	processor.ProcessOrder(order)

	// Process order using International Shipping
	internationalShipping := InternationalShipping{Country: "Germany"}
	processor = NewOrderProcessor(internationalShipping)
	processor.ProcessOrder(order)
}
