package services

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {

	if product, err := NewSage().GetProductDetail("COTDRILLDK"); err == nil {
		if product.Success {
			fmt.Printf("Received Quantity of %02.2f\n", product.Response.QtyInStock)
		} else {
			fmt.Println("Product not found")
		}
	} else {
		fmt.Printf("Received error %v\n", err)
	}

}
