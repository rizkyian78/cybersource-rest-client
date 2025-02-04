// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/rizkyian78/cybersource-rest-client/client/asymmetric_key_management"
	"github.com/rizkyian78/cybersource-rest-client/client/capture"
	"github.com/rizkyian78/cybersource-rest-client/client/chargeback_details"
	"github.com/rizkyian78/cybersource-rest-client/client/chargeback_summaries"
	"github.com/rizkyian78/cybersource-rest-client/client/conversion_details"
	"github.com/rizkyian78/cybersource-rest-client/client/credit"
	"github.com/rizkyian78/cybersource-rest-client/client/customer"
	"github.com/rizkyian78/cybersource-rest-client/client/customer_payment_instrument"
	"github.com/rizkyian78/cybersource-rest-client/client/customer_shipping_address"
	"github.com/rizkyian78/cybersource-rest-client/client/decision_manager"
	"github.com/rizkyian78/cybersource-rest-client/client/download_d_t_d"
	"github.com/rizkyian78/cybersource-rest-client/client/download_x_s_d"
	"github.com/rizkyian78/cybersource-rest-client/client/instrument_identifier"
	"github.com/rizkyian78/cybersource-rest-client/client/interchange_clearing_level_details"
	"github.com/rizkyian78/cybersource-rest-client/client/invoice_settings"
	"github.com/rizkyian78/cybersource-rest-client/client/invoices"
	"github.com/rizkyian78/cybersource-rest-client/client/key_generation"
	"github.com/rizkyian78/cybersource-rest-client/client/net_fundings"
	"github.com/rizkyian78/cybersource-rest-client/client/notification_of_changes"
	"github.com/rizkyian78/cybersource-rest-client/client/payer_authentication"
	"github.com/rizkyian78/cybersource-rest-client/client/payment_batch_summaries"
	"github.com/rizkyian78/cybersource-rest-client/client/payment_instrument"
	"github.com/rizkyian78/cybersource-rest-client/client/payments"
	"github.com/rizkyian78/cybersource-rest-client/client/payouts"
	"github.com/rizkyian78/cybersource-rest-client/client/purchase_and_refund_details"
	"github.com/rizkyian78/cybersource-rest-client/client/refund"
	"github.com/rizkyian78/cybersource-rest-client/client/report_definitions"
	"github.com/rizkyian78/cybersource-rest-client/client/report_downloads"
	"github.com/rizkyian78/cybersource-rest-client/client/report_subscriptions"
	"github.com/rizkyian78/cybersource-rest-client/client/reports"
	"github.com/rizkyian78/cybersource-rest-client/client/retrieval_details"
	"github.com/rizkyian78/cybersource-rest-client/client/retrieval_summaries"
	"github.com/rizkyian78/cybersource-rest-client/client/reversal"
	"github.com/rizkyian78/cybersource-rest-client/client/search_transactions"
	"github.com/rizkyian78/cybersource-rest-client/client/secure_file_share"
	"github.com/rizkyian78/cybersource-rest-client/client/symmetric_key_management"
	"github.com/rizkyian78/cybersource-rest-client/client/taxes"
	"github.com/rizkyian78/cybersource-rest-client/client/tokenization"
	"github.com/rizkyian78/cybersource-rest-client/client/transaction_batches"
	"github.com/rizkyian78/cybersource-rest-client/client/transaction_details"
	"github.com/rizkyian78/cybersource-rest-client/client/user_management"
	"github.com/rizkyian78/cybersource-rest-client/client/user_management_search"
	"github.com/rizkyian78/cybersource-rest-client/client/verification"
	"github.com/rizkyian78/cybersource-rest-client/client/void"
)

// Default cybersource HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "apitest.cybersource.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new cybersource HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cybersource {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cybersource HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cybersource {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cybersource client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cybersource {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cybersource)
	cli.Transport = transport
	cli.AsymmetricKeyManagement = asymmetric_key_management.New(transport, formats)
	cli.Capture = capture.New(transport, formats)
	cli.ChargebackDetails = chargeback_details.New(transport, formats)
	cli.ChargebackSummaries = chargeback_summaries.New(transport, formats)
	cli.ConversionDetails = conversion_details.New(transport, formats)
	cli.Credit = credit.New(transport, formats)
	cli.Customer = customer.New(transport, formats)
	cli.CustomerPaymentInstrument = customer_payment_instrument.New(transport, formats)
	cli.CustomerShippingAddress = customer_shipping_address.New(transport, formats)
	cli.DecisionManager = decision_manager.New(transport, formats)
	cli.Downloaddtd = download_d_t_d.New(transport, formats)
	cli.Downloadxsd = download_x_s_d.New(transport, formats)
	cli.InstrumentIdentifier = instrument_identifier.New(transport, formats)
	cli.InterchangeClearingLevelDetails = interchange_clearing_level_details.New(transport, formats)
	cli.InvoiceSettings = invoice_settings.New(transport, formats)
	cli.Invoices = invoices.New(transport, formats)
	cli.KeyGeneration = key_generation.New(transport, formats)
	cli.NetFundings = net_fundings.New(transport, formats)
	cli.NotificationOfChanges = notification_of_changes.New(transport, formats)
	cli.PayerAuthentication = payer_authentication.New(transport, formats)
	cli.PaymentBatchSummaries = payment_batch_summaries.New(transport, formats)
	cli.PaymentInstrument = payment_instrument.New(transport, formats)
	cli.Payments = payments.New(transport, formats)
	cli.Payouts = payouts.New(transport, formats)
	cli.PurchaseAndRefundDetails = purchase_and_refund_details.New(transport, formats)
	cli.Refund = refund.New(transport, formats)
	cli.ReportDefinitions = report_definitions.New(transport, formats)
	cli.ReportDownloads = report_downloads.New(transport, formats)
	cli.ReportSubscriptions = report_subscriptions.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.RetrievalDetails = retrieval_details.New(transport, formats)
	cli.RetrievalSummaries = retrieval_summaries.New(transport, formats)
	cli.Reversal = reversal.New(transport, formats)
	cli.SearchTransactions = search_transactions.New(transport, formats)
	cli.SecureFileShare = secure_file_share.New(transport, formats)
	cli.SymmetricKeyManagement = symmetric_key_management.New(transport, formats)
	cli.Taxes = taxes.New(transport, formats)
	cli.Tokenization = tokenization.New(transport, formats)
	cli.TransactionBatches = transaction_batches.New(transport, formats)
	cli.TransactionDetails = transaction_details.New(transport, formats)
	cli.UserManagement = user_management.New(transport, formats)
	cli.UserManagementSearch = user_management_search.New(transport, formats)
	cli.Verification = verification.New(transport, formats)
	cli.Void = void.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Cybersource is a client for cybersource
type Cybersource struct {
	AsymmetricKeyManagement asymmetric_key_management.ClientService

	Capture capture.ClientService

	ChargebackDetails chargeback_details.ClientService

	ChargebackSummaries chargeback_summaries.ClientService

	ConversionDetails conversion_details.ClientService

	Credit credit.ClientService

	Customer customer.ClientService

	CustomerPaymentInstrument customer_payment_instrument.ClientService

	CustomerShippingAddress customer_shipping_address.ClientService

	DecisionManager decision_manager.ClientService

	Downloaddtd download_d_t_d.ClientService

	Downloadxsd download_x_s_d.ClientService

	InstrumentIdentifier instrument_identifier.ClientService

	InterchangeClearingLevelDetails interchange_clearing_level_details.ClientService

	InvoiceSettings invoice_settings.ClientService

	Invoices invoices.ClientService

	KeyGeneration key_generation.ClientService

	NetFundings net_fundings.ClientService

	NotificationOfChanges notification_of_changes.ClientService

	PayerAuthentication payer_authentication.ClientService

	PaymentBatchSummaries payment_batch_summaries.ClientService

	PaymentInstrument payment_instrument.ClientService

	Payments payments.ClientService

	Payouts payouts.ClientService

	PurchaseAndRefundDetails purchase_and_refund_details.ClientService

	Refund refund.ClientService

	ReportDefinitions report_definitions.ClientService

	ReportDownloads report_downloads.ClientService

	ReportSubscriptions report_subscriptions.ClientService

	Reports reports.ClientService

	RetrievalDetails retrieval_details.ClientService

	RetrievalSummaries retrieval_summaries.ClientService

	Reversal reversal.ClientService

	SearchTransactions search_transactions.ClientService

	SecureFileShare secure_file_share.ClientService

	SymmetricKeyManagement symmetric_key_management.ClientService

	Taxes taxes.ClientService

	Tokenization tokenization.ClientService

	TransactionBatches transaction_batches.ClientService

	TransactionDetails transaction_details.ClientService

	UserManagement user_management.ClientService

	UserManagementSearch user_management_search.ClientService

	Verification verification.ClientService

	Void void.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cybersource) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AsymmetricKeyManagement.SetTransport(transport)
	c.Capture.SetTransport(transport)
	c.ChargebackDetails.SetTransport(transport)
	c.ChargebackSummaries.SetTransport(transport)
	c.ConversionDetails.SetTransport(transport)
	c.Credit.SetTransport(transport)
	c.Customer.SetTransport(transport)
	c.CustomerPaymentInstrument.SetTransport(transport)
	c.CustomerShippingAddress.SetTransport(transport)
	c.DecisionManager.SetTransport(transport)
	c.Downloaddtd.SetTransport(transport)
	c.Downloadxsd.SetTransport(transport)
	c.InstrumentIdentifier.SetTransport(transport)
	c.InterchangeClearingLevelDetails.SetTransport(transport)
	c.InvoiceSettings.SetTransport(transport)
	c.Invoices.SetTransport(transport)
	c.KeyGeneration.SetTransport(transport)
	c.NetFundings.SetTransport(transport)
	c.NotificationOfChanges.SetTransport(transport)
	c.PayerAuthentication.SetTransport(transport)
	c.PaymentBatchSummaries.SetTransport(transport)
	c.PaymentInstrument.SetTransport(transport)
	c.Payments.SetTransport(transport)
	c.Payouts.SetTransport(transport)
	c.PurchaseAndRefundDetails.SetTransport(transport)
	c.Refund.SetTransport(transport)
	c.ReportDefinitions.SetTransport(transport)
	c.ReportDownloads.SetTransport(transport)
	c.ReportSubscriptions.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.RetrievalDetails.SetTransport(transport)
	c.RetrievalSummaries.SetTransport(transport)
	c.Reversal.SetTransport(transport)
	c.SearchTransactions.SetTransport(transport)
	c.SecureFileShare.SetTransport(transport)
	c.SymmetricKeyManagement.SetTransport(transport)
	c.Taxes.SetTransport(transport)
	c.Tokenization.SetTransport(transport)
	c.TransactionBatches.SetTransport(transport)
	c.TransactionDetails.SetTransport(transport)
	c.UserManagement.SetTransport(transport)
	c.UserManagementSearch.SetTransport(transport)
	c.Verification.SetTransport(transport)
	c.Void.SetTransport(transport)
}
