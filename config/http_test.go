package config

import (
	"net/http"
	"testing"
)

// Tests the digest production against the sample code at: https://developer.cybersource.com/api/developer-guides/dita-gettingstarted/authentication/GenerateHeader/httpSignatureAuthentication.html
func TestDigest(t *testing.T) {

	{
		digest, err := generateDigest([]byte("sample body"))
		if err != nil {
			t.Fatal(err)
		}

		if expect := "SHA-256=Pk8hyTs9x7soFEhGvFNUrXL4wcwIXouMvxCgRCicswk="; digest != expect {
			t.Fatalf("Expected '%s' but got '%s'", expect, digest)
		}
	}

	{
		digest, err := generateDigest([]byte(body))
		if err != nil {
			t.Fatal(err)
		}

		if expect := "SHA-256=/BpZpLkjYFrp2vhJ1VzcoLWpyMkOJY+ZA97FRN69SGM="; digest != expect {
			t.Fatalf("Expected '%s' but got '%s'", expect, digest)
		}
	}

}

const body = `{"clientReferenceInformation": {"code": "TC50171_3" },
  "processingInformation": { "commerceIndicator": "internet"
  },
  "paymentInformation": {
    "card": {"number": "4111111111111111", "expirationMonth": "12",
      "expirationYear": "2031",
      "securityCode": "123"
    }
  },
  "orderInformation": {
    "amountDetails": {
      "totalAmount": "102.21",
      "currency": "USD"
    },
    "billTo": {
      "firstName": "John",
      "lastName": "Doe",
      "address1": "1 Market St",
      "address2": "Address 2",
      "locality": "san francisco",
      "administrativeArea": "CA",
      "postalCode": "94105",
      "country": "US",
      "email": "test@cybs.com",
      "phoneNumber": "4158880000"
    }
  }
}`

// Test the signature production against the website at: https://developer.cybersource.com/api-reference-assets/index.html#payments
func TestGenerateSignatureValue(t *testing.T) {

	// keyid="08c94330-f618-42a3-b09d-e1e43be5efda", algorithm="HmacSHA256", headers="host (request-target) digest v-c-merchant-id", signature="HBEajAuU2SR+ajr+Y++TBwJjBhEflZo6u0tvbNwF77Y="

	key := "yBJxy6LjM2TmcPGu+GaJrHtkke25fPpUX+UY6/L/1tE="

	expect := `HBEajAuU2SR+ajr+Y++TBwJjBhEflZo6u0tvbNwF77Y=`

	req := &httpRequest{
		headers: http.Header{
			"v-c-merchant-id": []string{"testrest"},
			"host":            []string{"apitest.cybersource.com"},
			"digest":          []string{"SHA-256=E9P5jdXRZjM5aGk6JGDxiiV5bEiWz5pBRhvWzHw+TXQ="},
		},
		body:        []byte(`{"clientReferenceInformation":{"code":"TC50171_3"},"merchantDefinedInformation":null,"orderInformation":{"amountDetails":{"amexAdditionalAmounts":null,"currency":"USD","taxDetails":null,"totalAmount":"13.37"},"lineItems":null},"paymentInformation":{"card":{"expirationMonth":"12","expirationYear":"2031","number":"4111111111111111","securityCode":"123"}},"processingInformation":{"commerceIndicator":"internet"}}`),
		path:        "/pts/v2/payments/",
		method:      "POST",
		queryParams: nil,
	}

	headerNames := []string{"host", "(request-target)", "digest", "v-c-merchant-id"}

	signature, err := generateHTTPSignatureValue(key, req, headerNames...)
	if err != nil {
		t.Fatal(err)
	}

	if expect != signature {
		t.Fatalf("Expected signature '%s' but found '%s'!", expect, signature)
	}

}
