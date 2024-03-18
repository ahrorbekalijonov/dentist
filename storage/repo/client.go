package repo

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

type GetAllClient struct {
	Page  int
	Limit int
}

type NewClientI interface {
	CreateClient(*Client) (*Client, error)
	GetClient(id string) (*Client, error)
	UpdateClient(*Client) (*Client, error)
	DeleteClient(id string) (bool, error)
	GetAllClients(*GetAllClient) (*AllClients, error)
	GetAllClientsCount() (int, error)
	SearchClients(str string) (*AllClients, error)
}
