package models

// Ticket : Representation of the tickets table from the database
type Ticket struct {
	IDTicket      *int64   `json:"idTicket,omitempty"`
	Date          *string  `json:"date,omitempty"`
	Amount        *float64 `json:"amount,omitempty"`
	File          *string  `json:"file,omitempty"`
	Currency      *string  `json:"currency,omitempty"`
	PaymentMethod *string  `json:"paymentMethod,omitempty"`
	IDReport      *int64   `json:"idReport,omitempty"`
	IDConcept     *int64   `json:"idConcept,omitempty"`
	IDDeleted     *bool    `json:"isDeleted,omitempty"`

	Report  *Report  `json:"report,omitempty"`
	Concept *Concept `json:"concept,omitempty"`
}
