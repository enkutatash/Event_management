package repository

import (
	"context"
	"database/sql"
	"errors"
	"event/models"
)

type AdminRepository interface {
	CreateEvent(*models.Event) (*string, error)
	CancelEvent(*string) error
}

type adminRepo struct {
	db *sql.DB
}

// CancelEvent implements AdminRepository.
func (a *adminRepo) CancelEvent(eventId *string) error {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	query := "DELETE FROM events WHERE id = $1"
	res, err := a.db.ExecContext(c, query, eventId)
	if err != nil {
		return errors.New("failed to cancel event")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("failed to retrieve affected rows")
	}
	if rowsAffected == 0 {
		return errors.New("event not found")
	}

	return nil
}

// CreateEvent implements AdminRepository.
func (a *adminRepo) CreateEvent(event *models.Event) (*string, error) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	query := "INSERT INTO events(title, description, location, start_date, end_date, start_time, end_time, price, quota, organizer) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	var eventId string
	query2 := "INSERT INTO event_quotas(event_id, remaining_quota) VALUES($1, $2)"
	err := a.db.QueryRowContext(c, query, event.Title, event.Description, event.Location, event.StartDate, event.EndDate, event.StartTime, event.EndTime, event.Price, event.Quota, event.Organizer).Scan(&eventId)
	if err != nil {
		return nil, err
	}
	_, err = a.db.ExecContext(c, query2, eventId, event.Quota)
	if err != nil {
		return nil, err
	}
	return &eventId, nil
}

func NewAdminRepo(db *sql.DB) AdminRepository {
	return &adminRepo{
		db: db,
	}
}
