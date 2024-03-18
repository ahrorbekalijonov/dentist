package models

type Appointment struct {
	Id string
	ClientId string
	Date string
	Diagnostics string
	Treatment string
	Amount int
}

type ReqAppointment struct {
	ClientId string
	Date string
	Diagnostics string
	Treatment string
	Amount int
}

type New struct {
	ClientName string
	PhoneNumber string
	ClientId string
	AppointmentId string
	Date string
	Diagnostics string
	Treatment string
	Amount int
}

type ReqNew struct {
	ClientName string
	PhoneNumber string
	Date string
	Diagnostics string
	Treatment string
	Amount int
}

type AllAppointments struct {
	Appointments []*Appointment
}
