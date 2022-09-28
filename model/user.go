package model

import "time"

type User struct {
	ID        int64  `json:"id" gorm:"type:int primary key auto_increment"`
	Name      string `json:"name" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255)"`
	Role      string `json:"role" gorm:"type:varchar(255)"`
	Password  string
	CreatedAt time.Time `json:"created_at"`
}
