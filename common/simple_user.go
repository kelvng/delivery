package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first-name" gorm:"column:first_name;"`
	Role      string `json:"role" gorm:"column:role;"`
	Avatar    string `json:"avatar,omitempty" gorm:"column:avatar; type:json;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
