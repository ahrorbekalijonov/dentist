package repo

type Appointment struct {
	Id string
	ClientId string
	Date string
	Diagnostics string
	Treatment string
	Amount int
}

type AllAppointments struct {
	Appointment []*Appointment
}

type GetAllAppointment struct{
	Page int
	Limit int
}

type NewAppointmentI interface {
	CreateAppointment(*Appointment)(*Appointment, error)
	GetAppointment(id string)(*Appointment, error)
	UpdateAppointment(*Appointment)(*Appointment, error)
	DeleteAppointment(id string)(bool, error)
	GetAllAppointments(*GetAllAppointment)(*AllAppointments, error)
	GetAppointmentsWithDate(req, page, limit int) (*AllAppointments, error)
	GetAppointmentsWithClientId(id string, page, limit int) ([]Appointment, error)
}