package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"helpstack/models"
	"helpstack/pkg"
)

// UserHandler .
type UserHandler struct {
	UsrRepo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{repo}
}

func (uh *UserHandler) GetUsers(c *fiber.Ctx) error {

	usr, err:=uh.UsrRepo.AllUsers(c.Context(), models.Filter{
		Limit: 0,
		Page:  0,
	})
	if err!=nil{
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(usr)
}

func (uh *UserHandler) GetOneUser(c *fiber.Ctx) error{
	id:= c.Params("id")
	//contxt := c.Context()
	//var ctx, _ = context.WithTimeout(contxt, 30*time.Second)
	user, err := uh.UsrRepo.UserById(c.Context(), id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(user)

}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error{

	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return pkg.BadRequest("Invalid params")
	}
	usr, err:=uh.UsrRepo.StoreUser(c.Context(), user)
	if err != nil {
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(usr)

}



