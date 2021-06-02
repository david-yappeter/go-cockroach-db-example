package model

type User struct {
	ID   int `json:"id" gorm:"type:SERIAL PRIMARY KEY;"`
	Name string `json:"name" gorm:"type:varchar(255) not null"`
}
