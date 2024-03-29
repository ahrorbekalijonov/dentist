package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/dentist/storage/repo"
	"github.com/jmoiron/sqlx"
)

type appoinmentRepo struct {
	db *sqlx.DB
}

func NewAppointmentRepo(db *sqlx.DB) repo.NewAppointmentI {
	return &appoinmentRepo{
		db: db,
	}
}

//This method create a new appointment
func (h *appoinmentRepo) CreateAppointment(req *repo.Appointment) (*repo.Appointment, error) {
	query := `
	INSERT INTO 
		appointments(
			id,
			client_id,
			date,
			diagnostics,
			treatment,
			amount
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, client_id, date, diagnostics, treatment, amount
	`
	var nullTime sql.NullTime
	tx, err := h.db.Begin()
	if err != nil {
		log.Println("Error creating transaction create appointment: ", err)
		return nil, err
	}
	var user repo.Appointment
	err = tx.QueryRow(
		query,
		req.Id,
		req.ClientId,
		req.Date,
		req.Diagnostics,
		req.Treatment,
		req.Amount,
	).Scan(
		&user.Id,
		&user.ClientId,
		&nullTime,
		&user.Diagnostics,
		&user.Treatment,
		&user.Amount,
	)
	if err != nil {
		log.Println("Error to create appointment in database: ", err)
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if nullTime.Valid {
		user.Date = nullTime.Time.Format("2006-01-02 15:04:05")
	}

	return &user, nil
}

//This method delete appointment with id
func (h *appoinmentRepo) DeleteAppointment(id string) (bool, error) {
	query := `
	UPDATE 
	    appointments
	SET 
	    deleted_at = CURRENT_TIMESTAMP
	WHERE 
		id = $1
	AND
	    deleted_at IS NULL`

	tx, err := h.db.Begin()
	if err != nil {
		log.Println("Error creating transaction to delete appointment")
		return false, err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		log.Println("Error to deleting appointment in database: ", err)
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

//This method get appointment with id
func (h *appoinmentRepo) GetAppointment(id string) (*repo.Appointment, error) {
	query := `
	SELECT 
	    id,
		client_id,
        date,
        diagnostics,
        treatment,
		amount
	FROM 
	    appointments
	WHERE 
		id = $1
	AND 
	    deleted_at IS NULL`
	var appointment repo.Appointment
	err := h.db.QueryRow(query, id).Scan(
		&appointment.Id,
		&appointment.ClientId,
		&appointment.Date,
		&appointment.Diagnostics,
		&appointment.Treatment,
		&appointment.Amount,
	)
	if err != nil {
		log.Println("Error to get appointment in database: ", err)
		return nil, err
	}

	return &appointment, nil
}

//This method update appointment with id
func (h *appoinmentRepo) UpdateAppointment(req *repo.Appointment) (*repo.Appointment, error) {
	query := `
	UPDATE 
		appointments
	SET
		client_id = $1,
		date = $2,
        diagnostics = $3,
        treatment = $4,
        amount = $5
	WHERE
		id = $6
	AND 
		deleted_at IS NULL
	RETURNING id, client_id, date, diagnostics, treatment, amount`
	var user repo.Appointment
	err := h.db.QueryRow(
		query,
		req.ClientId,
		req.Date,
		req.Diagnostics,
		req.Treatment,
		req.Amount,
		req.Id,
	).Scan(
		&user.Id,
		&user.ClientId,
		&user.Date,
		&user.Diagnostics,
		&user.Treatment,
		&user.Amount,
	)
	if err != nil {
		log.Println("Error updating appointment in database: ", err)
	}

	return &user, nil
}

//This method get all appointments with page and limit
func (h *appoinmentRepo) GetAllAppointments(req *repo.GetAllAppointment) (*repo.AllAppointments, error) {
	query := `
	SELECT
		id,
		client_id,
		date,
		diagnostics,
        treatment,
        amount
	FROM 
	    appointments
	WHERE 
	    deleted_at IS NULL
	ORDER BY date
	LIMIT $1
	OFFSET $2`
	offset := req.Limit * (req.Page - 1)
	rows, err := h.db.Query(query, req.Limit, offset)
	if err != nil {
		log.Println("Error to get all appointments in database: ", err)
		return nil, err
	}
	var AllAppointments = repo.AllAppointments{}
	for rows.Next() {
		var appointment repo.Appointment
		err = rows.Scan(
			&appointment.Id,
			&appointment.ClientId,
			&appointment.Date,
			&appointment.Diagnostics,
			&appointment.Treatment,
			&appointment.Amount,
		)
		if err != nil {
			log.Println("Error to get all appointments: ", err)
			return nil, err
		}
		defer rows.Close()
		AllAppointments.Appointment = append(AllAppointments.Appointment, &appointment)
	}

	return &AllAppointments, nil
}

//this method takes appointments by number, that is, if a positive number is entered, 
// it will take the data of that number of days from today, if it is negative, 
// it will take the data of the previous day
func (h *appoinmentRepo) GetAppointmentsWithDate(req, page, limit int) (*repo.AllAppointments, error) {
	now := time.Now().Format("2006-01-02")
	to := time.Now().AddDate(0, 0, req).Format("2006-01-02")
	if req < 0 {
		now, to = to, now
	}

	offset := limit * (page - 1)
	query := `
	SELECT 
		id,
		client_id,
		date,
		diagnostics,
		treatment,
		amount
	FROM 
		appointments
	WHERE 
		date >= $1 AND date < $2 AND deleted_at IS NULL
	ORDER BY date
	LIMIT $3
	OFFSET $4`

	rows, err := h.db.Query(query, now, to, limit, offset)
	if err != nil {
		log.Println("Error get appointments with date", err)
		return nil, err
	}
	var appointments repo.AllAppointments
	for rows.Next() {
		var appointment repo.Appointment
		err = rows.Scan(
			&appointment.Id,
			&appointment.ClientId,
			&appointment.Date,
			&appointment.Diagnostics,
			&appointment.Treatment,
			&appointment.Amount,
		)
		if err != nil {
			log.Println("Error get appointments with date", err)
			return nil, err
		}
		appointments.Appointment = append(appointments.Appointment, &appointment)
		defer rows.Close()
	}

	return &appointments, nil
}

//This method get appointments with client id
func (h *appoinmentRepo) GetAppointmentsWithClientId(id string, page, limit int) ([]repo.Appointment, error) {
	query := `
	SELECT 
		id,
		client_id,
		date,
		diagnostics,
		treatment,
		amount
	FROM 
		appointments
	WHERE
		client_id = $1
	AND 
		deleted_at IS NULL 
	ORDER BY date
	LIMIT $2
	OFFSET $3 `
	offset := limit * (page - 1)
	rows, err := h.db.Query(query, id, limit, offset)
	if err != nil {
		log.Println("Error to get appointment with course_id", err)
		return nil, err
	}

	var appointments []repo.Appointment
	for rows.Next() {
		var appointment repo.Appointment
		err = rows.Scan(
			&appointment.Id,
			&appointment.ClientId,
			&appointment.Date,
			&appointment.Diagnostics,
			&appointment.Treatment,
			&appointment.Amount,
		)
		if err != nil {
			log.Println("Error to get appointment with course_id", err)
			return nil, err
		}
		defer rows.Close()
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
