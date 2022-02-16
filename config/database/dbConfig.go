package dbPkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"strings"
)


type Databases struct {
	Gorm *gorm.DB
}


type SqlDbConfig struct {
	Driver string
	Host string
	Username string
	Password string
	Port   int
	DbName string
}

//type Database struct {
//	*gorm.DB
//}


func NewSqlDb(config *SqlDbConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	var confg gorm.Config
	if os.Getenv("ENABLE_GORM_LOGGER") != "" {
		confg = gorm.Config{}
	} else {
		confg = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}
	switch strings.ToLower(config.Driver) {
	case "mysql":
		dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DbName + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		db, err = gorm.Open(mysql.Open(dsn), &confg)
		break
	case "postgresql", "postgres":
		//dsn := "user=" + config.Username + " password=" + config.Image + " dbname=" + config.DbName + " host=" + config.Host + " port=" + strconv.Itoa(config.ServerPort) + " TimeZone=UTC"
		dsn := fmt.Sprintf("user=%s dbname=%s password=%s  sslmode=disable", config.Username, config.DbName, config.Password)
		//dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", config.Username, config.DbName, config.Image, config.Host)
		db, err = gorm.Open(postgres.Open(dsn), &confg)
		break
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) + "?database=" + config.DbName
		db, err = gorm.Open(sqlserver.Open(dsn), &confg)
		break
	}
	if err != nil {
		log.Fatal(err)
		//panic("Failed to connect database")
	}
	fmt.Println("Connection Opened to DbName")
	return db, err
}

