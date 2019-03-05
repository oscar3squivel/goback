package models

// Concept : Representation of the concepts table from the database
type Concept struct {
	IDConcept   *int64    `json:"idConcept,omitempty"`
	Description *string   `json:"description,omitempty"`
	Notes       *string   `json:"notes,omitempty"`
	IsDeleted   *bool     `json:"isDeleted,omitempty"`
	Tickets     *[]Ticket `json:"tickets,omitempty"`
}
