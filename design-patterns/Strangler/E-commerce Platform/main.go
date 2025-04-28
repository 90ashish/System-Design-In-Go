package main

import (
	"fmt"
	"net/http"
)

// LegacySystem - Simulates the old system
type LegacySystem struct{}

func (l *LegacySystem) GetProductDetails(productID string) string {
	return fmt.Sprintf("[Legacy System] Product %s details retrieved.", productID)
}

// NewSystem - Simulates the new microservices-based system
type NewSystem struct{}

func (n *NewSystem) GetProductDetails(productID string) string {
	return fmt.Sprintf("[New System] Product %s details retrieved with modern architecture.", productID)
}

// Router - The Strangler Pattern's key element
type Router struct {
	legacySystem *LegacySystem
	newSystem    *NewSystem
}

// NewRouter - Initializes the Router with both systems
func NewRouter() *Router {
	return &Router{
		legacySystem: &LegacySystem{},
		newSystem:    &NewSystem{},
	}
}

// GetProductDetails - Directs requests to the correct system
func (r *Router) GetProductDetails(productID string) string {
	// Assume product IDs greater than "5000" are handled by the new system
	if productID >= "5000" {
		return r.newSystem.GetProductDetails(productID)
	}
	return r.legacySystem.GetProductDetails(productID)
}

// HTTP Handler - Demonstrating API Routing
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	productID := req.URL.Query().Get("productID")
	if productID == "" {
		http.Error(w, "Product ID is required", http.StatusBadRequest)
		return
	}

	response := r.GetProductDetails(productID)
	fmt.Fprintln(w, response)
}

func main() {
	router := NewRouter()

	// Simulating HTTP server for routing requests
	http.Handle("/product", router)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
