package main

import (
	"log"

	// "github.com/go-openapi/strfmt"
	"github.com/tooolbox/cybersource-rest-client-go/client"
)

func main() {

	c := client.Default
	created, err := c.Payments.CreatePayment()

	log.Printf("Created: %#v", created)
	log.Printf("Err: %v", err)
}
