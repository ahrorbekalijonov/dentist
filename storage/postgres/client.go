package postgres

import (
	"log"

	"github.com/dentist/storage/repo"
	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) repo.NewClientI {
	return &clientRepo{
		db: db,
	}
}

// This function is create a client
func (h *clientRepo) CreateClient(c *repo.Client) (*repo.Client, error) {
	query := `
	INSERT INTO
		clients(
			id,
		    name,
            last_name,
            father_name,
            phone_number,
            address,
			birth_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, name, last_name, father_name, phone_number, address, birth_date`
	tx, err := h.db.Begin()
	if err != nil {
		log.Println("Error creating transaction create client: ", err)
		return nil, err
	}
	var user repo.Client
	err = tx.QueryRow(
		query,
		c.Id,
		c.Name,
		c.LastName,
		c.FatherName,
		c.PhoneNumber,
		c.Address,
		c.BirthDate,
	).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.FatherName,
		&user.PhoneNumber,
		&user.Address,
		&user.BirthDate,
	)
	if err != nil {
		log.Println("Error to creating client in database: ", err)
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// This function is delete a client with client id
func (h *clientRepo) DeleteClient(id string) (bool, error) {
	query := `
	UPDATE 
		clients
	SET 
		deleted_at = CURRENT_TIMESTAMP
	WHERE 
		id = $1
	AND
		deleted_at IS NULL`
	tx, err := h.db.Begin()
	if err != nil {
		log.Println("Error creating transaction delete client: ", err)
		return false, err
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		log.Println("Error to delete client in database: ", err)
		tx.Rollback()
		return false, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return false, err
	}

	return true, nil
}

// This function is get a client with client id
func (h *clientRepo) GetClient(id string) (*repo.Client, error) {
	query := `
	SELECT 
	    id,
        name,
        last_name,
        father_name,
        phone_number,
		address,
        birth_date
	FROM 
	    clients
	WHERE 
		id = $1
	AND 
	    deleted_at IS NULL`
	tx, err := h.db.Begin()
	if err != nil {
		log.Println("Error creating transaction get client: ", err)
		return nil, err
	}
	var user repo.Client
	err = tx.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.FatherName,
		&user.PhoneNumber,
		&user.Address,
		&user.BirthDate,
	)
	if err != nil {
		log.Println("Error to get client in database: ", err)
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// This function is update a client with client id
func (h *clientRepo) UpdateClient(c *repo.Client) (*repo.Client, error) {
	query := `
	UPDATE 
		clients
	SET
		name = $1,
		last_name = $2,
		father_name = $3,
        phone_number = $4,
        address = $5,
        birth_date = $6
	WHERE 
	    id = $7
	AND 
	    deleted_at IS NULL
	RETURNING
		id, name, last_name, father_name, phone_number, address, birth_date`

	tx, err := h.db.Begin()
	if err != nil {
		log.Println("Error creating transaction update client: ", err)
		return nil, err
	}
	var user repo.Client
	err = tx.QueryRow(
		query,
		c.Name,
		c.LastName,
		c.FatherName,
		c.PhoneNumber,
		c.Address,
		c.BirthDate,
		c.Id,
	).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.FatherName,
		&user.PhoneNumber,
		&user.Address,
		&user.BirthDate,
	)
	if err != nil {
		log.Println("Error to updating client: ", err)
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}

// This function is get all clients with given page and limit
func (h *clientRepo) GetAllClients(req *repo.GetAllClient) (*repo.AllClients, error) {
	query := `
	SELECT
		id,
		name,
        last_name,
        father_name,
        phone_number,
        address,
		birth_date
	FROM 
		clients
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`
	offset := req.Limit * (req.Page - 1)
	rows, err := h.db.Query(query, req.Limit, offset)
	if err != nil {
		log.Println("Error to get all clients: ", err)
		return nil, err
	}
	defer rows.Close()
	var clients repo.AllClients
	for rows.Next() {
		var client repo.Client
		err = rows.Scan(
			&client.Id,
			&client.Name,
			&client.LastName,
			&client.FatherName,
			&client.PhoneNumber,
			&client.Address,
			&client.BirthDate,
		)
		if err != nil {
			log.Println("Error to get all clients: ", err)
			return nil, err
		}
		clients.Clients = append(clients.Clients, &client)
	}

	return &clients, nil
}

// This function is get all clients count
func (h *clientRepo) GetAllClientsCount() (int, error) {
	query := `SELECT COUNT(*) FROM clients WHERE deleted_at IS NULL`
	var resp int
	err := h.db.QueryRow(query).Scan(&resp)
	if err != nil {
		log.Println("Error to get all clients count")
		return 0, err
	}

	return resp, nil
}

// This function is searching clients with name or last_name
func (h *clientRepo) SearchClients(str string) (*repo.AllClients, error) {
	query := `
	SELECT 
		id, 
		name, 
		last_name, 
		father_name, 
		phone_number, 
		address, 
		birth_date
	FROM 
		clients
	WHERE 
		deleted_at IS NULL 
	AND
		name 
	ILIKE 
		'%` + str + `%' OR last_name ILIKE '%` + str + `%'`

	var Clients repo.AllClients
	rows, err := h.db.Query(query)
	if err != nil {
		log.Println("Error search clients in database", err)
		return nil, err
	}

	for rows.Next() {
		var client repo.Client
		err = rows.Scan(
			&client.Id,
			&client.Name,
			&client.LastName,
			&client.FatherName,
			&client.PhoneNumber,
			&client.Address,
			&client.BirthDate,
		)
		if err != nil {
			log.Println("Error search clients in database", err)
			return nil, err
		}
		Clients.Clients = append(Clients.Clients, &client)
	}

	return &Clients, nil
}
