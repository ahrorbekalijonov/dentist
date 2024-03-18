package storage

import (
	"github.com/dentist/storage/postgres"
	"github.com/dentist/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Client() repo.NewClientI
	Appointment() repo.NewAppointmentI
}

type storagePg struct {
	clientRepo repo.NewClientI
	appoinmentRepo repo.NewAppointmentI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
        clientRepo: postgres.NewClientRepo(db),
        appoinmentRepo: postgres.NewAppointmentRepo(db),
    }
}

func (s *storagePg) Appointment() repo.NewAppointmentI {
	return s.appoinmentRepo
}
func (s *storagePg) Client() repo.NewClientI {
	return s.clientRepo
}