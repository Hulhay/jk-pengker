package store

type StoreListResponse struct {
	ID       int64    `json:"id" gorm:"type:int primary key auto_increment"`
	Name     string   `json:"name" gorm:"type:varchar(255)"`
	Category []string `json:"category"`
}
