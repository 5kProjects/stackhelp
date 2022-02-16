package user

import (
	"context"
	"gorm.io/gorm"
	"helpstack/models"
)

// NewUserGormRepo will create a new object of
func NewUserGormRepo(db *gorm.DB) UserRepository {
	return &UserGormRepo{conn: db}
}


// UserRepository specifies application user related database operations
type UserRepository interface {
	AllUsers(ctx context.Context, filter models.Filter) ([]User, error)
	UserById(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *User) (int64, int64, error)
	DeleteUser(ctx context.Context, id string) (int64, error)
	StoreUser(ctx context.Context, user *User) (*User, error)
}

// UserGormRepo implements the item.ItemRepository interface
type UserGormRepo struct {
	conn *gorm.DB
}

func (u UserGormRepo) AllUsers(ctx context.Context, filter models.Filter) ([]User, error) {
	var usrs []User

	err := u.conn.Find(&usrs).Error
	if err!=nil{
		return nil, err
	}
	return usrs, err

}

func (u UserGormRepo) UserById(ctx context.Context, id string) (*User, error) {
	user:= User{}
	errs:=u.conn.First(&user).Error

	if errs!=nil{
		return nil, errs
	}
	return &user, nil
}

func (u UserGormRepo) UpdateUser(ctx context.Context, user *User) (int64, int64, error) {
	panic("implement me")
}

func (u UserGormRepo) DeleteUser(ctx context.Context, id string) (int64, error) {
	panic("implement me")
}

func (u UserGormRepo) StoreUser(ctx context.Context, user *User) (*User, error) {

	err := u.conn.Create(user).Error
	if err!=nil {
		return nil, err
	}
	return user, err
}



