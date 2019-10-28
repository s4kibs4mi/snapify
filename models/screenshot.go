package models

import "time"

type Screenshot struct {
	ID         string    `json:"id" sql:"id" gorm:"primary_key"`
	Website    string    `json:"website" sql:"website" gorm:"index"`
	StoredPath string    `json:"stored_path" sql:"stored_path" gorm:"index"`
	CreatedAt  time.Time `json:"created_at" sql:"created_at" gorm:"index"`
	UpdatedAt  time.Time `json:"updated_at" sql:"updated_at" gorm:"index"`
}

func (ss *Screenshot) TableName() string {
	return "screenshots"
}
