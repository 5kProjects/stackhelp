package models

import (
	"gorm.io/gorm"
	"time"
)

type DefaultModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

// User model
type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" query:"name"`
	Password string `json:"-" form:"-" query:"-"`
	Email    string
	RoleID   uint `gorm:"column:role_id" json:"role_id"`
	Role     Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

// Role model
type Role struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name" form:"name" query:"name"`
	Description string `gorm:"type:varchar(100);" json:"description"`
}

// Role model
type Tags struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name" form:"name" query:"name"`
	Description string `gorm:"type:varchar(100);" json:"description"`
}

// Article model
type Article struct {
	gorm.Model
	Name string `json:"name" form:"name" query:"name"`

	Body  string
	Title string

	UserId uint `gorm:"column:user_id" json:"user_id"`
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type Filter struct{
	Limit int
	Page int
}

type Question struct {
	Body      string
	RefId     string
	Upvote    []string
	DownVotes []string
	UserId    uint `gorm:"column:user_id" json:"user_id"`
	User      User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

type Answer struct {
	Body       string
	QuestionId string `gorm:"column:question_id" json:"question_id"`
	Upvote     []string
	DownVotes  []string

	UserId   uint `gorm:"column:user_id" json:"user_id"`
	User     User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Question User `gorm:"foreignKey:QuestionId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
