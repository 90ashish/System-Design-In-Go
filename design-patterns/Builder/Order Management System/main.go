package main

import "fmt"

// Order struct - Immutable once created
type Order struct {
	orderID         string
	shippingAddress string
	billingAddress  string
	giftWrap        bool
	discountCode    string
	orderNotes      string
}

// OrderBuilder - Builder Class
type OrderBuilder struct {
	orderID         string
	shippingAddress string
	billingAddress  string
	giftWrap        bool
	discountCode    string
	orderNotes      string
}

// NewOrderBuilder - Initializes a new OrderBuilder instance
func NewOrderBuilder(orderID string) *OrderBuilder {
	return &OrderBuilder{
		orderID: orderID,
	}
}

// WithShippingAddress - Adds a shipping address
func (b *OrderBuilder) WithShippingAddress(address string) *OrderBuilder {
	b.shippingAddress = address
	return b
}

// WithBillingAddress - Adds a billing address
func (b *OrderBuilder) WithBillingAddress(address string) *OrderBuilder {
	b.billingAddress = address
	return b
}

// WithGiftWrap - Adds gift wrapping
func (b *OrderBuilder) WithGiftWrap() *OrderBuilder {
	b.giftWrap = true
	return b
}

// WithDiscountCode - Adds a discount code
func (b *OrderBuilder) WithDiscountCode(code string) *OrderBuilder {
	b.discountCode = code
	return b
}

// WithOrderNotes - Adds order notes
func (b *OrderBuilder) WithOrderNotes(notes string) *OrderBuilder {
	b.orderNotes = notes
	return b
}

// Build - Final step to construct the immutable Order object
func (b *OrderBuilder) Build() Order {
	return Order{
		orderID:         b.orderID,
		shippingAddress: b.shippingAddress,
		billingAddress:  b.billingAddress,
		giftWrap:        b.giftWrap,
		discountCode:    b.discountCode,
		orderNotes:      b.orderNotes,
	}
}

// DisplayOrder - Displays the order details
func DisplayOrder(order Order) {
	fmt.Println("\n🛒 Order Details:")
	fmt.Printf("Order ID: %s\n", order.orderID)
	fmt.Printf("Shipping Address: %s\n", order.shippingAddress)
	fmt.Printf("Billing Address: %s\n", order.billingAddress)
	fmt.Printf("Gift Wrap: %v\n", order.giftWrap)
	fmt.Printf("Discount Code: %s\n", order.discountCode)
	fmt.Printf("Order Notes: %s\n", order.orderNotes)
}

func main() {
	// Order 1: Full Customization
	order1 := NewOrderBuilder("ORD123").
		WithShippingAddress("123 Amazon St, Seattle, WA").
		WithBillingAddress("456 Prime Ave, Seattle, WA").
		WithGiftWrap().
		WithDiscountCode("NEWYEAR50").
		WithOrderNotes("Handle with care.").
		Build()

	DisplayOrder(order1)

	// Order 2: Basic Order with Minimal Configurations
	order2 := NewOrderBuilder("ORD124").
		WithShippingAddress("789 Cloud Drive, Austin, TX").
		Build()

	DisplayOrder(order2)
}
