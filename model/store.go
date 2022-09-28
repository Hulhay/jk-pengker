package model

import "time"

type Store struct {
	ID        int64     `json:"id" gorm:"type:int primary key auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Owner     string    `json:"owner" gorm:"type:varchar(255)"`
	Category  string    `json:"category"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
}
