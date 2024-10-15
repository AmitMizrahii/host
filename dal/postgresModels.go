package dal

type UserModel struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

func (UserModel) TableName() string {
	return "users"
}
