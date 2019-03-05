package models

// Client : Representation of the clients table from the database
type Client struct {
	IDClient  *int64    `json:"idClient,omitempty"`
	Name      *string   `json:"name,omitempty"`
	Address   *string   `json:"address,omitempty"`
	IsDeleted *bool     `json:"isDeleted,omitempty"`
	Reports   *[]Report `json:"reports,omitempty"`
}
