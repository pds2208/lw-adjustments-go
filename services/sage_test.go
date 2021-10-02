package services

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: time.RFC3339})

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
