package api

import "time"

type Reservation struct {
	ID       uint          `json:"id" validate:"required"`
	Date     time.Time     `json:"date" validate:"required"`
	Duration time.Duration `json:"duration" validate:"required"`
}

type Room struct {
	Name        string `json:"name" validate:"required"`
	Number      string `json:"number" validate:"required"`
	SeatsNumber int    `json:"seatsNumber" validate:"required"`
}

type User struct {
	Name    string            `json:"name" validate:"required"`
	Contact map[string]string `json:"contact" validate:"required"`
}

type Notification struct {
	Reservation Reservation `json:"reservation" validate:"required"`
	Room        Room        `json:"room" validate:"required"`
	User        User        `json:"user" validate:"required"`
}
