package tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"helpstack/config"
	dbPkg "helpstack/config/database"
	"helpstack/pkg/users"
	"helpstack/server"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func TestMain(m *testing.M) {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf:=config.New()
	Db, err:= dbPkg.NewSqlDb(&conf.SqlDbConfig)

	app = server.Create()
	Db.AutoMigrate(&user.User{})


	usrH:= user.NewUserHandler(user.NewUserGormRepo(Db))
	v1 := app.Group("/api/v1")
	server.SetUserRoutes(v1, *usrH)
	//Routes(app)

	// Cleanup books
	Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&user.User{})

	exitVal := m.Run()

	// Cleanup books
	Db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&user.User{})

	os.Exit(exitVal)
}


func TestCreateUser(t *testing.T) {
	newUser := map[string]interface{}{
		"name":  "abebe",
		"email": "abebe@gmail.com",
		"password": "abc123",
	}
	body, _ := json.Marshal(newUser)
	req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	body, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode, "status ok")

	var usr user.User
	err := json.Unmarshal(body, &usr)
	assert.Equal(t, err, nil)

	assert.NotEqual(t, usr.ID, nil)
	assert.Equal(t, usr.Name, newUser["name"])
	assert.Equal(t, usr.Email, newUser["email"])
	assert.Equal(t, usr.Password, newUser["rating"])
}
