package models

import "time"

type Client struct {
	Name    string
	Sitting bool
}

type Event struct {
	Time time.Time
	ID   int
	Body string
}

type Table struct {
	Number      int
	Revenue     int
	TotalTime   time.Duration
	CurrentTime time.Time
	IsOccupied  bool
	OccupiedBy  *Client
}
