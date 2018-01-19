package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Country Model
type Country struct {
	ID        string     `sql:"size:3" gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Customers []Customer `gorm:"ForeignKey:CountryID;AssociationForeignKey:ID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// Seed Countries table
func (c Country) Seed(db *gorm.DB) {
	countries := []Country{
		{ID: "MX", Name: "Mexico"},
	}

	for _, country := range countries {
		db.Create(&country)
	}
}
