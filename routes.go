package qrtickets

import "net/http"

// Route - Define information necessary to route a url request to handler function
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	AdminOnly bool
}

var routes = []Route{{
	"",
	"GET",
	"/api/v1/events",
	EventList,
	false,
}, {
	"EventShow",
	"GET",
	"/api/v1/events/{eventId}",
	EventShow,
	false,
}, {
	"GenSignature",
	"GET",
	"/gensig",
	GenSignature,
	true,
}, {
	"VerifySig",
	"GET",
	"/api/v1/tickets/{sig1}/{sig2}/{hash}",
	VerifySignature,
	true,
}, {
	"GenerateTicket",
	"GET",
	"/api/v1/tickets/generate/{hash}",
	GenTicket,
	true,
}, {
	"LoadConf",
	"GET",
	"/loadconf",
	WebConfLoad,
	true,
}, {
	"TestSign",
	"GET",
	"/testsign",
	TestSign,
	true,
}}
