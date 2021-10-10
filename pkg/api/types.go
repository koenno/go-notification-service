package api

import "time"

type Reservation struct {
	ID       uint          `json:"id"`
	Date     time.Time     `json:"date"`
	Duration time.Duration `json:"duration"`
}

type Room struct {
	Name        string `json:"name"`
	Number      string `json:"number"`
	SeatsNumber int    `json:"seatsNumber"`
}

type User struct {
	Name    string            `json:"name"`
	Contact map[string]string `json:"contact"`
}

type Notification struct {
	Reservation Reservation `json:"reservation"`
	Room        Room        `json:"room"`
	User        User        `json:"user"`
}
