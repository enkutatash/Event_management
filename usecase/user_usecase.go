package usecase

import (
	"errors"
	"event/models"
	"event/repository"
)

type UserUsecase interface {
	GetAllEvents(offset int,limit int) (*[]models.EventRes, error)
	GetEventById(eventId *string) (*models.EventRes, error)
	BookTicket(eventId *string,userId *int,tickerNo *int) (error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// BookTicket implements UserUsecase.
func (u *userUsecase) BookTicket(eventId *string,userId *int,tickerNo *int) error{
	// check eventId
	_,err :=u.userRepo.GetEventById(eventId)
	if err != nil{
		return err
	}

	//check available tickers
	available,err := u.userRepo.AvailableTicket(eventId)
	if available < *tickerNo{
		return errors.New("no enough ticket")
	}
	err  = u.userRepo.BookTicket(eventId,userId,tickerNo)
	if err != nil{
		return err
	}
	return nil
	//book ticket
}

// GetAllEvents implements UserUsecase.
func (u *userUsecase) GetAllEvents(offset int,limit int) (*[]models.EventRes, error) {
	events, err := u.userRepo.GetAllEvents(offset,limit)
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
