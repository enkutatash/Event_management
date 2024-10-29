package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"event/models"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type UserRepository interface {
	GetAllEvents() ([]models.EventRes, error)
	GetEventById(*string) (*models.EventRes, error)
	BookTicker(eventId *string, userId *int) error
}

type userRepository struct {
	db *sql.DB
	cache *redis.Client
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
	value, err := u.cache.Get(c, query).Result()

	if err == redis.Nil {
		fmt.Println("from db")
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

		eventJSON, err := json.Marshal(events)
		if err != nil {
			return nil,err
		}

		err = u.cache.Set(c, query,eventJSON, time.Minute*15).Err()
		if err != nil {
			return nil, err
		}
		return events, nil
	}else if err != nil {
		return nil, err
	}else{
		fmt.Println("from cache")
		var events [] models.EventRes
		err = json.Unmarshal([]byte(value), &events)
		if err != nil {
			return nil, err
		}
		return events, nil
	}

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

func NewUserRepo(db *sql.DB,cache *redis.Client) UserRepository {
	return &userRepository{
		db: db,
		cache: cache,
	}
}
