package article

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"helpstack/models"
	"helpstack/pkg"
)

// ArticleHandler .
type ArticleHandler struct {
	UsrRepo ArticleRepository
}

func NewUserHandler(repo ArticleRepository) *ArticleHandler {
	return &ArticleHandler{repo}
}

func (uh *ArticleHandler) GetMany(c *fiber.Ctx) error {

	usr, err:=uh.UsrRepo.FetchMany(c.Context(), models.Filter{
		Limit: 0,
		Page:  0,
	})
	if err!=nil{
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(usr)
}

func (uh *ArticleHandler) GetOne(c *fiber.Ctx) error{
	id:= c.Params("id")
	//contxt := c.Context()
	//var ctx, _ = context.WithTimeout(contxt, 30*time.Second)
	user, err := uh.UsrRepo.GetOneById(c.Context(), id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(user)

}

func (uh *ArticleHandler) CreateOne(c *fiber.Ctx) error{
	user := new(Article)
	if err := c.BodyParser(user); err != nil {
		return pkg.BadRequest("Invalid params")
	}
	usr, err:=uh.UsrRepo.CreateOne(c.Context(), user)
	if err != nil {
		return pkg.Unexpected(err.Error())
	}
	return c.JSON(usr)

}



