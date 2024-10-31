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
	GetAllEvents(offset int,limit int) ([]models.EventRes, error)
	GetEventById(*string) (*models.EventRes, error)
	BookTicket(eventId *string, userId *int, tickerNo *int) error
	AvailableTicket(eventId *string) (int, error)
}

type userRepository struct {
	db    *sql.DB
	cache *redis.Client
}

// AvailableTicket implements UserRepository.
func (u *userRepository) AvailableTicket(eventId *string) (int, error) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	query := "SELECT remaining_quota FROM event_quotas WHERE event_id = $1"
	var quota int
	err := u.db.QueryRowContext(c, query, eventId).Scan(&quota)
	if err != nil {
		return 0, err
	}
	return quota, nil
}

// BookTicker implements UserRepository.
func (u *userRepository) BookTicket(eventId *string, userId *int, ticketNo *int) error {
    c, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Start a transaction
    // tx, err := u.db.BeginTx(c, nil)
    // if err != nil {
    //     return errors.New("failed to begin transaction")
    // }

    // Prepare the queries
    insertQuery := "INSERT INTO bookings(event_id, user_id, number_of_tickets) VALUES($1, $2, $3)"
    updateQuery := "UPDATE event_quotas SET remaining_quota = remaining_quota - $1 WHERE event_id = $2"

    // Execute the insert query
    _, err := u.db.ExecContext(c, insertQuery, eventId, userId, ticketNo)
    if err != nil {
         // Rollback transaction on error
        return errors.New("failed to book ticket")
    }

    // Execute the update query
    _, err = u.db.ExecContext(c, updateQuery, ticketNo, eventId)
    if err != nil {
       // Rollback transaction on error
        return errors.New("failed to update remaining quota")
    }

    // Commit the transaction
   

    return nil
}


// GetAllEvents implements UserRepository.
func (u *userRepository) GetAllEvents(offset int, limit int) ([]models.EventRes, error) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	var events []models.EventRes
	cacheKey := fmt.Sprintf("events:limit:%d:offset:%d", limit, offset)
	query := "SELECT id, title, description, location, start_date, end_date, start_time, end_time, price, organizer FROM events LIMIT $1 OFFSET $2"
	value, err := u.cache.Get(c, cacheKey).Result()

	if err == redis.Nil {
		fmt.Println("from db")
		rows, err := u.db.QueryContext(c, query, limit, offset)
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
			return nil, err
		}

		err = u.cache.Set(c, cacheKey, eventJSON, time.Minute*15).Err()
		if err != nil {
			return nil, err
		}
		return events, nil
	} else if err != nil {
		return nil, err
	} else {
		fmt.Println("from cache")
		var events []models.EventRes
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

func NewUserRepo(db *sql.DB, cache *redis.Client) UserRepository {
	return &userRepository{
		db:    db,
		cache: cache,
	}
}
