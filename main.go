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

	req := payments.NewCreatePaymentParams().WithCreatePaymentRequest(payments.CreatePaymentBody{
		ClientReferenceInformation: &payments.CreatePaymentParamsBodyClientReferenceInformation{
			Code: "TC50171_3",
		},
		ProcessingInformation: &payments.CreatePaymentParamsBodyProcessingInformation{
			CommerceIndicator: "internet",
		},
		PaymentInformation: &payments.CreatePaymentParamsBodyPaymentInformation{
			Card: &payments.CreatePaymentParamsBodyPaymentInformationCard{
				Number:          "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2031",
				SecurityCode:    "123",
			},
		},
		OrderInformation: &payments.CreatePaymentParamsBodyOrderInformation{
			AmountDetails: &payments.CreatePaymentParamsBodyOrderInformationAmountDetails{
				TotalAmount: "13.37",
				Currency:    "USD",
			},
		},
	})

	created, err := c.Payments.CreatePayment(req)

	log.Printf("\n\n")
	log.Printf("Created: %v\n", created)
	log.Printf("Err: %v\n", err)
}
