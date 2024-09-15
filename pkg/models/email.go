package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/niharikabhavaraju/go_scheduler/pkg/config"
)

var db *gorm.DB

type Email struct {
	gorm.Model
	Subject string    `json:"subject"`
	Body    string    `json:"body"`
	Status  string    `json:"status"`
	To      string    `json:"to"`
	Time    time.Time `json:"time"`
}

// InitDB initializes the database connection
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Email{})
}

func (e *Email) CreateEmail() (*Email, error) {
	isPrimarykey := db.NewRecord(e)
	if !isPrimarykey {
		return nil, errors.New("no primary key")
	}
	res := db.Create(&e)
	return e, res.Error
}

func GetEmails() ([]Email, error) {
	var Emails []Email
	res := db.Find(&Emails)
	return Emails, res.Error
}

func GetEmailById(ID int) (*Email, *gorm.DB, error) {
	var Email Email
	res := db.First(&Email, ID)
	return &Email, db, res.Error
}

func DeleteEmail(ID int) (Email, error) {
	var email Email
	res := db.Where("ID = ?", ID).Delete(email)
	return email, res.Error
}

func GetEmailByTime(startTime, endTime string) ([]Email, error) {
	var Emails []Email
	res := db.Where("time BETWEEN ? AND ?", startTime, endTime).Find(&Emails)
	return Emails, res.Error
}

func UpdateEmailStatus(email *Email, status string) (*Email, error) {
	res := db.Model(&email).Update("status", status)
	return email, res.Error
}
