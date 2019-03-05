package models

// Employee : Representation of the employees table from the database
type Employee struct {
	IDEmployee      *int64    `json:"idEmployee,omitempty"`
	Email           *string   `json:"email,omitempty"`
	Name            *string   `json:"name,omitempty"`
	PhoneNumber     *string   `json:"phoneNumber,omitempty"`
	BirthDate       *string   `json:"birthDate,omitempty"`
	Nationality     *string   `json:"nationality,omitempty"`
	EmergencyNumber *string   `json:"emergencyNumber,omitempty"`
	Password        *string   `json:"password,omitempty"`
	Role            *string   `json:"role,omitempty"`
	IsDeleted       *bool     `json:"isDeleted,omitempty"`
	Reports         *[]Report `json:"reports,omitempty"`
}
