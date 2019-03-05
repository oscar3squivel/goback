package models

import "time"

// Report : Representation of the reports table from the database
// ReportStatus : { A: 'Accepted', R: 'Rejected', P: 'Pending'}
type Report struct {
	IDReport             *int64     `json:"idReport,omitempty"`
	OriginPlace          *string    `json:"originPlace,omitempty"`
	DestinationPlace     *string    `json:"destinationPlace,omitempty"`
	DepartureDate        *time.Time `json:"departureDate,omitempty"`
	ReturnDate           *time.Time `json:"returnDate,omitempty"`
	TeamName             *string    `json:"teamName,omitempty"`
	IDLeader             *int64     `json:"idLeader,omitempty"`
	Purpose              *string    `json:"purpose,omitempty"`
	IsSponsored          *bool      `json:"isSponsored,omitempty"`
	ClientPurpose        *string    `json:"clientPurpose,omitempty"`
	SpecialRequirements  *string    `json:"specialRequirements,omitempty"`
	OperationDate        *time.Time `json:"operationDate,omitempty"`
	OperationArrangement *string    `json:"operationArrangement,omitempty"`
	StatusChangeDate     *string    `json:"statusChangeDate,omitempty"`
	Comments             *string    `json:"comments,omitempty"`
	ReportStatus         *string    `json:"reportStatus,omitempty"`
	IDClient             *int64     `json:"idClient,omitempty"`
	IDEmployee           *int       `json:"idEmployee,omitempty"`
	IsInternational      *bool      `json:"isInternational,omitempty"`
	IsDeleted            *bool      `json:"isDeleted,omitempty"`

	Tickets  *[]Ticket `json:"tickets,omitempty"`
	Employee *Employee `json:"employee,omitempty"`
	Client   *Client   `json:"client,omitempty"`
}
