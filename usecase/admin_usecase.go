package usecase

import (
	"event/models"
	"event/repository"

	
)



type AdminUsecase interface {
	CreateEvent(event *models.Event) (*string, error)
	CancelEvent(eventId *string) error
}

type adminUsecase struct {
	AdminRepo repository.AdminRepository
}

// CancelEvent implements AdminUsecase.
func (a *adminUsecase) CancelEvent(eventId *string) error {
	return a.AdminRepo.CancelEvent(eventId)
}

// CreateEvent implements AdminUsecase.
func (a *adminUsecase) CreateEvent(event *models.Event) (*string, error) {
	return a.AdminRepo.CreateEvent(event)
}

func NewAdminUsecase(repo repository.AdminRepository) AdminUsecase {
	return &adminUsecase{
		AdminRepo: repo,
	}
}
