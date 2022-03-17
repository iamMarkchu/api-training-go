package model

type CategoryModel struct {
	Id          uint64 `json:"id" gorm:"column:id"`
	ParentId    uint64 `json:"parent_id" gorm:"column:parent_id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Status      uint8  `json:"status" gorm:"column:status"`
	UserId      uint64 `json:"user_id" gorm:"column:user_id"`
	CreatedAt   int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   int64  `json:"updated_at" gorm:"column:updated_at"`
}

func (CategoryModel) TableName() string {
	return "iron_categories"
}
