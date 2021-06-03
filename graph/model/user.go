package model

type User struct {
	ID   int    `json:"id" gorm:"type:int8;PRIMARY KEY;default:nextVal('user_seq')"`
	Name string `json:"name" gorm:"type:varchar(255) not null"`
}
