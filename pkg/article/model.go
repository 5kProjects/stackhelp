package article

import (
	"gorm.io/gorm"
)

// Article model
type Article struct {
	gorm.Model
	Title string `json:"title" form:"title" query:"title"`
	Image string `json:"image" form:"image" query:"-"`
	Body  string `json:"body" form:"body" query:"-"`

}

