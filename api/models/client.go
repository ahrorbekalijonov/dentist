package models

import "github.com/dentist/storage/repo"

type Client struct {
	Id          string
	Name        string
	LastName    string
	FatherName  string
	PhoneNumber string
	Address     string
	BirthDate   string
	AllAppointments []repo.Appointment
}

type AllClients struct {
	Clients []*Client
}

type ReqClient struct {
	Name        string
	LastName    string
	FatherName  string
	PhoneNumber string
	Address     string
	BirthDate   string
}
