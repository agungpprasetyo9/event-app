package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserId      uint      `json:"userId"`
	User        User      `gorm:"foreignKey:UserId" json:"-"`
	DateTime    time.Time `json:"datetime" binding:"required"`
}

var events []Event = []Event{}

// fungsi simpan event
func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvent() []Event {
	return events
}
