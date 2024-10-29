package repository

import (
	"context"
	"database/sql"
	"errors"
	"event/models"
)

type UserRepository interface {
	GetAllEvents() ([]models.EventRes, error)
	GetEventById(*string) (*models.EventRes, error)
	BookTicker(eventId *string, userId *int) error
}

type userRepository struct {
	db *sql.DB
}

// BookTicker implements UserRepository.
func (u *userRepository) BookTicker(eventId *string, userId *int) error {
	panic("unimplemented")
}

// GetAllEvents implements UserRepository.
func (u *userRepository) GetAllEvents() ([]models.EventRes, error) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	var events []models.EventRes

	query := "SELECT id, title, description, location, start_date, end_date, start_time, end_time, price, organizer FROM events"

	rows, err := u.db.QueryContext(c, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.EventRes
		if err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.Location,
			&event.StartDate, &event.EndDate, &event.StartTime, &event.EndTime,
			&event.Price, &event.Organizer); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil

}

// GetEventById implements UserRepository.
func (u *userRepository) GetEventById(eventId *string) (*models.EventRes, error) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	query := "SELECT id, title, description, location, start_date, end_date, start_time, end_time, price, organizer FROM events WHERE id = $1"

	var event models.EventRes
	err := u.db.QueryRowContext(c, query, eventId).Scan(
		&event.Id,
		&event.Title,
		&event.Description,
		&event.Location,
		&event.StartDate,
		&event.EndDate,
		&event.StartTime,
		&event.EndTime,
		&event.Price,
		&event.Organizer,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("event not found")
		}
		return nil, err
	}

	return &event, nil 

}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
