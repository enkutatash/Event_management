package usecase

import (
	"event/models"
	"event/repository"
)

type UserUsecase interface {
	GetAllEvents() (*[]models.EventRes, error)
	GetEventById(eventId *string) (*models.EventRes, error)
	BookTicket()
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// BookTicket implements UserUsecase.
func (u *userUsecase) BookTicket() {
	panic("unimplemented")
}

// GetAllEvents implements UserUsecase.
func (u *userUsecase) GetAllEvents() (*[]models.EventRes, error) {
	events, err := u.userRepo.GetAllEvents()
	if err != nil {
		return nil, err
	}
	return &events, nil
}

// GetEventById implements UserUsecase.
func (u *userUsecase) GetEventById(eventid *string) (*models.EventRes, error) {
	return u.userRepo.GetEventById(eventid)
}

func NewuserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}
