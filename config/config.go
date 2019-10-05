package config

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

const (
	AuthTypeHTTPSignature = "http_signature"
	AuthTypeJWT           = "jwt"
)

type Config struct {
	AuthenticationType string
	RunEnvironment     string
	MerchantId         string

	// http_signature parameters
	MerchantKeyId     string
	MerchantSecretKey string

	// jwt parameters
	KeysDirectory string
	KeyFileName   string
	KeyAlias      string
	KeyPass       string
}

func (cfg Config) ClientAuthInfoWriter() (runtime.ClientAuthInfoWriter, error) {
	switch cfg.AuthenticationType {
	case AuthTypeHTTPSignature:
		return cfg.HTTPSignatureAuth(), nil
	case AuthTypeJWT:
		return nil, fmt.Errorf("not implemented")
	default:
		return nil, fmt.Errorf("Unknown AuthenticationType '%s', please use one of: [%s, %s]", cfg.AuthenticationType, AuthTypeHTTPSignature, AuthTypeJWT)
	}
}

type AuthenticateRequestFunc func(req runtime.ClientRequest, reg strfmt.Registry) error

func (f AuthenticateRequestFunc) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	return f(req, reg)
}
