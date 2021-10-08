package api

import "time"

type Reservation struct {
	ID       uint
	Date     time.Time
	Duration time.Duration
}

type Room struct {
	Name        string
	Number      string
	SeatsNumber int
}

type User struct {
	Name    string
	Contact map[string]string
}

type Notification struct {
	Reservation Reservation
	Room        Room
	User        User
}
