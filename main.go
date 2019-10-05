package main

import (
	"log"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tooolbox/cybersource-rest-client-go/client"
	"github.com/tooolbox/cybersource-rest-client-go/client/payments"
)

func main() {

	tr := httptransport.New(
		client.DefaultHost,
		client.DefaultBasePath,
		client.DefaultSchemes)
	tr.Consumers["application/json;charset=utf-8"] = runtime.JSONConsumer()
	tr.Producers["application/json;charset=utf-8"] = runtime.JSONProducer()
	// tr.DefaultAuthentication =

	c := client.New(tr, strfmt.Default)
	created, err := c.Payments.CreatePayment(payments.NewCreatePaymentParams())

	log.Printf("Created: %#v", created)
	log.Printf("Err: %v", err)
}
