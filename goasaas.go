package goasaas

import (
	"fmt"

	"github.com/nwiry/goasaas/customers"
	"github.com/nwiry/goasaas/request"
)

const (
	ENVIRONMENT_PRODUCTION = "https://www.asaas.com/" // Production environment URL for Asaas API.
	ENVIRONMENT_SANDBOX    = "https://sandbox.asaas.com/" // Sandbox environment URL for Asaas API.
	API_VERSION            = "v3" // Version of the Asaas API being used.
)

// Asaas is the main client struct for interacting with the Asaas API.
type Asaas struct {
	access_token string   // The access token used for authentication with the Asaas API.
	environment  string   // The environment URL for making API requests (production or sandbox).
}

// request returns a new instance of the Request type with the API endpoint set based on the Asaas environment.
func (a *Asaas) request() *request.Request {
	r := &request.Request{}
	e := fmt.Sprintf("%s/api/%s", a.environment, API_VERSION)
	r.SetEndpoint(e)
	return r
}

// New creates a new instance of the Asaas client with the provided access token and environment.
// The environment can be either ENVIRONMENT_PRODUCTION or ENVIRONMENT_SANDBOX.
func New(access_token, enviroment string) *Asaas {
	if enviroment != ENVIRONMENT_PRODUCTION && ENVIRONMENT_SANDBOX != enviroment {
		panic("invalid asaas environment")
	}

	return &Asaas{
		access_token: access_token, // The access token used for authentication with the Asaas API.
		environment:  enviroment,   // The environment URL for making API requests (production or sandbox).
	}
}

// The following functions in the Asaas struct are placeholders for handling various Asaas resources, but they are not implemented yet.

// Accounts is a placeholder function for handling Asaas accounts, but it's not implemented yet.
func (a *Asaas) Accounts() {

}

// Anticipations is a placeholder function for handling Asaas anticipations, but it's not implemented yet.
func (a *Asaas) Anticipations() {

}

// Bill is a placeholder function for handling Asaas bills, but it's not implemented yet.
func (a *Asaas) Bill() {

}

// Customers returns a new instance of the Customers type, which is used to interact with Asaas customers.
func (a *Asaas) Customers() *customers.Customers {
	return &customers.Customers{Request: a.request()}
}

// Installments is a placeholder function for handling Asaas installments, but it's not implemented yet.
func (a *Asaas) Installments() {

}

// MobilePhoneRecharges is a placeholder function for handling Asaas mobile phone recharges, but it's not implemented yet.
func (a *Asaas) MobilePhoneRecharges() {

}

// Notifications is a placeholder function for handling Asaas notifications, but it's not implemented yet.
func (a *Asaas) Notifications() {

}

// PaymentLinks is a placeholder function for handling Asaas payment links, but it's not implemented yet.
func (a *Asaas) PaymentLinks() {

}

// Payments is a placeholder function for handling Asaas payments, but it's not implemented yet.
func (a *Asaas) Payments() {

}

// Pix is a placeholder function for handling Asaas Pix, but it's not implemented yet.
func (a *Asaas) Pix() {

}

// Serasa is a placeholder function for handling Asaas Serasa, but it's not implemented yet.
func (a *Asaas) Serasa() {

}

// Subscriptions is a placeholder function for handling Asaas subscriptions, but it's not implemented yet.
func (a *Asaas) Subscriptions() {

}

// Transfers is a placeholder function for handling Asaas transfers, but it's not implemented yet.
func (a *Asaas) Transfers() {

}

// Webhook is a placeholder function for handling Asaas webhooks, but it's not implemented yet.
func (a *Asaas) Webhook() {

}
