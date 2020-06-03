// Code generated by go-swagger; DO NOT EDIT.

package download_x_s_d

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new download x s d API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for download x s d API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetXSDV2(params *GetXSDV2Params) (*GetXSDV2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetXSDV2 useds to download x s ds for reports

  Downloads XSDs for reports on no-auth.
*/
func (a *Client) GetXSDV2(params *GetXSDV2Params) (*GetXSDV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetXSDV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getXSDV2",
		Method:             "GET",
		PathPattern:        "/xsds/{reportDefinitionNameVersion}",
		ProducesMediaTypes: []string{"text/xml"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetXSDV2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetXSDV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getXSDV2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
