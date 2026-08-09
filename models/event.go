package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	UserId      int    `json:"userId"`
}

var events []Event = []Event{}

// fungsi simpan event
func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvent() []Event {
	return events
}
