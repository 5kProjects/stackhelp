package server

import (
	"github.com/gofiber/fiber/v2"
	"helpstack/config/database"
	"helpstack/pkg/article"
	"helpstack/pkg/users"
)

func SetupApiRoutes(app *fiber.App, dbs dbPkg.Databases) {

	usrH:= user.NewUserHandler(user.NewUserGormRepo(dbs.Gorm))
	articleH:= article.NewUserHandler(article.NewArticleGormRepo(dbs.Gorm))
	v1 := app.Group("/api/v1")
	SetUserRoutes(v1, *usrH)
	SetArticleRoutes(v1, *articleH)
}


func SetUserRoutes(route fiber.Router, usrH user.UserHandler) {
	route.Get("/user", usrH.GetUsers)
	route.Get("/user/:id", usrH.GetOneUser)
	//route.Put("/books/:id", UpdateBook)
	route.Post("/user", usrH.CreateUser)
	//route.Delete("/books/:id", DeleteBook)
}


func SetArticleRoutes(route fiber.Router, aH article.ArticleHandler) {
	route.Get("/article", aH.GetMany)
	route.Get("/articles", aH.GetFunction)
	route.Get("/article/:id", aH.GetOne)
	//route.Put("/article/:id", UpdateBook)
	route.Post("/article", aH.CreateOne)
	//route.Delete("/article/:id", DeleteBook)
}


