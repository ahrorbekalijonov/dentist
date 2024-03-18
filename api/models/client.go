package models

type Client struct {
	Id          string
	Name        string
	LastName    string
	FatherName  string
	PhoneNumber string
	Address     string
	BirthDate   string
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
