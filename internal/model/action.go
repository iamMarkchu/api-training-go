package model

type ActionModel struct {
	Id          uint64 `json:"id" gorm:"column:id"`
	CateId      uint64 `json:"cate_id" gorm:"column:cate_id"`
	Name        string `json:"name"  gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Status      uint8  `json:"status" gorm:"column:status"`
	UserId      uint64 `json:"user_id" gorm:"column:user_id"`
	CreatedAt   int64  `json:"created_at"  gorm:"column:created_at"`
	UpdatedAt   int64  `json:"updated_at" gorm:"column:updated_at"`
}

func (ActionModel) TableName() string {
	return "iron_actions"
}
