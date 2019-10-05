package main

import (
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/tooolbox/cybersource-rest-client-go/client/payments"
	"github.com/tooolbox/cybersource-rest-client-go/config"
)

func main() {

	cfg := config.Config{
		AuthenticationType: config.AuthTypeHTTPSignature,
		RunEnvironment:     config.RunEnvSandbox,
		MerchantId:         "abc",
		MerchantKeyId:      "123",
		MerchantSecretKey:  "abc123",
	}

	c, err := cfg.NewHTTPClient(strfmt.Default)
	if err != nil {
		log.Fatal(err)
	}

	created, err := c.Payments.CreatePayment(payments.NewCreatePaymentParams())

	log.Printf("Created: %#v", created)
	log.Printf("Err: %#v", err)
}
