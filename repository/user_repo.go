package repository

import "event/models"

type UserRepository interface {
	GetAllEvents() ([]models.Event, error)
	GetEventById(*string) (*models.Event, error)
}