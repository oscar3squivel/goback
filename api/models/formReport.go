package models

type Form struct {
	IDReport             *int64
	OriginPlace          *string
	DestinationPlace     *string
	DepartureDate        *string
	ReturnDate           *string
	TeamName             *string
	IDLeader             *int64
	Purpose              *string
	IsSponsored          *bool
	ClientPurpose        *string
	SpecialRequirements  *string
	OperationDate        *string
	OperationArrangement *string
	StatusChangeDate     *string
	Comments             *string
	ReportStatus         *string
	IDClient             *int64
	IDEmployee           *int
	IsInternational      *bool
	IsDeleted            *bool

	Tickets  *[]Ticket
	Employee *Employee
	Client   *Client
}
