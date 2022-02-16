package user

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"password" form:"password" query:"-"`
	Email    string  `json:"email" form:"email" query:"-"`
	//Roles    pq.StringArray `gorm:"type:text[]"`
	//RoleID   uint        `gorm:"column:role_id" json:"role_id"`
	//Role     models.Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

