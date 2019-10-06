package config

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tooolbox/cybersource-rest-client-go/client"
)

const (
	AuthTypeHTTPSignature = "http_signature"
	AuthTypeJWT           = "jwt"
	RunEnvSandbox         = "cybersource.environment.sandbox"
	RunEnvProduction      = "cybersource.environment.production"
	ApiHostSandbox        = client.DefaultHost
	ApiHostProduction     = "api.cybersource.com"
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

	// logging
	DebugLogging bool
}

func (cfg Config) NewHTTPClient(formats strfmt.Registry) (*client.Cybersource, error) {
	tr, err := cfg.Transport()
	if err != nil {
		return nil, err
	}

	return client.New(tr, formats), nil
}

func (cfg Config) Transport() (*httptransport.Runtime, error) {

	var host string
	switch strings.ToLower(cfg.RunEnvironment) {
	case RunEnvSandbox:
		host = ApiHostSandbox
	case RunEnvProduction:
		host = ApiHostProduction
	default:
		return nil, fmt.Errorf("unknown run environment '%s', please use one of: [%s, %s]", cfg.RunEnvironment, RunEnvSandbox, RunEnvProduction)
	}

	tr := httptransport.New(
		host,
		client.DefaultBasePath,
		client.DefaultSchemes,
	)

	tr.Consumers["application/json;charset=utf-8"] = runtime.JSONConsumer()
	tr.Producers["application/json;charset=utf-8"] = runtime.JSONProducer()
	tr.SetDebug(cfg.DebugLogging)

	auth, err := cfg.ClientAuthInfoWriter()
	if err != nil {
		return nil, err
	}

	tr.DefaultAuthentication = auth

	return tr, nil
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
