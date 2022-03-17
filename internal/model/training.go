package model

type TrainingModel struct {
	Id          uint64 `json:"id" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	Img         string `json:"img" gorm:"column:img"`
	UserId      uint64 `json:"user_id" gorm:"column:user_id"`
	Status      uint8  `json:"status" gorm:"column:status"`
	CreatedAt   int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   int64  `json:"updated_at" gorm:"column:updated_at"`
}

func (TrainingModel) TableName() string {
	return "iron_trainings"
}

type TrainingItemModel struct {
	Id         uint64 `json:"id,omitempty" gorm:"column:id"`
	GroupId    uint64 `json:"group_id,omitempty" gorm:"column:group_id"`
	TrainingId uint64 `json:"training_id,omitempty" gorm:"column:training_id"`
	ActionId   uint64 `json:"action_id,omitempty" gorm:"column:action_id"`
	CountType  uint8  `json:"count_type,omitempty" gorm:"column:count_type"`
	Weight     uint32 `json:"weight,omitempty" gorm:"column:weight"`
	CountNum   uint32 `json:"count_num,omitempty" gorm:"column:count_num"`
	UserId     uint64 `json:"user_id,omitempty" gorm:"user_id"`
	Status     uint8  `json:"status,omitempty" gorm:"status"`
	CreatedAt  int64  `json:"created_at,omitempty" gorm:"created_at"`
	UpdatedAt  int64  `json:"updated_at,omitempty" gorm:"updated_at"`
}

func (TrainingItemModel) TableName() string {
	return "iron_training_item"
}
