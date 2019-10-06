package main

import (
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/tooolbox/cybersource-rest-client-go/client"
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
		DebugLogging:       true,
	}

	tr, err := cfg.Transport()
	if err != nil {
		log.Fatal(err)
	}

	c := client.New(tr, strfmt.Default)

	created, err := c.Payments.CreatePayment(payments.NewCreatePaymentParams())

	log.Printf("\n\n")
	log.Printf("Created: %v\n", created)
	log.Printf("Err: %v\n", err)
}
