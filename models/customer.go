package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"gitlab.com/lscortesc/go-gin/db"
	"gitlab.com/lscortesc/go-gin/forms"
	"golang.org/x/crypto/bcrypt"
)

// Customer Model
type Customer struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Country   Country   `gorm:"ForeignKey:CountryID;AssociationForeignKey:ID" json:"-"`
	CountryID string    `sql:"size:3" json:"country_id"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MigrateRelationships Customer
func (c Customer) MigrateRelationships(db *gorm.DB) {
	db.Model(c).AddUniqueIndex("idx_customer_email", "email")
	db.Model(c).AddForeignKey("country_id", "countries(id)", "RESTRICT", "RESTRICT")
}

// Register Customer
func (c Customer) Register(data forms.RegisterForm) (Customer, error) {
	c.Firstname = data.Firstname
	c.Lastname = data.Lastname
	c.Email = data.Email
	c.CountryID = data.CountryID

	password, _ := HashPassword(data.Password)
	c.Password = string(password)

	if err := db.GetConnection().Create(&c).Error; err != nil {
		return c, err
	}

	return c, nil
}

// HashPassword to Customer Register
func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes), err
}
