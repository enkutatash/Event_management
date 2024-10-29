package models

import "time"

type LoginRes struct {
	Id          int
	FullName    string
	Email       string
	AccessToken string
}

type UserReq struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	Id       int
	FullName string
	Email    string
	Password string
}

type Event struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Price       float64	`json:"price"`
	Quota       int `json:"quota"`
	Organizer   string `json:"organizer"`
}

type EventRes struct{
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Price       float64	`json:"price"`
	Organizer   string `json:"organizer"`
}