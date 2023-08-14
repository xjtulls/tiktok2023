package model

type TableUser struct {
	ID              uint64 `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"backgroundImage"`
}

// TableName 修改表名映射
func (user TableUser) TableName() string {
	return "users"
}
