package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // this field will not show in JSON if the value is empty or null
	Runtime   Runtime   `json:"runtime,omitempty"` // this field will not show in JSON if the value is empty or null
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"`
}
