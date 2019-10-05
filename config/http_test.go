package config

import (
	"testing"
)

// Tests the digest production against the sample code at: https://developer.cybersource.com/api/developer-guides/dita-gettingstarted/authentication/GenerateHeader/httpSignatureAuthentication.html
func TestDigest(t *testing.T) {

	digest, err := generateDigest([]byte("sample body"))
	if err != nil {
		t.Fatal(err)
	}

	if expect := "SHA-256=Pk8hyTs9x7soFEhGvFNUrXL4wcwIXouMvxCgRCicswk="; digest != expect {
		t.Fatalf("Expected '%s' but got '%s'", expect, digest)
	}

}
