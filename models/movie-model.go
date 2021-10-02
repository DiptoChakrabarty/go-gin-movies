package models

import (
	"fmt"
	"os"

	"github.com/DiptoChakrabarty/go-gin-movies/user"

	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type MovieModel interface {
	Add(movie operation.Movie)
	Update(movie operation.Movie)
	Delete(movie operation.Movie)
	GetAll() []operation.Movie
	GetOne(id uint64) operation.Movie
	AddUser(newuser user.User)
	GetUser(username string) user.User
	GetAllUsers() []user.User
}

type model struct {
	DBConn *gorm.DB
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getDB() (db *gorm.DB, err error) {
	db_type := getEnv("DB_TYPE", "sqlite3")
	db_connection_string := getEnv("DB_CONNECTION_STRING", "./db/movie.db")
	return gorm.Open(db_type, db_connection_string)
}

func NewModelDB() MovieModel {
	db, err := getDB()
	if err != nil {
		fmt.Println(err.Error())
		panic("Unable to connect to DB")
	}
	db.AutoMigrate(&operation.Movie{}, &operation.Actor{}, &user.User{})
	return &model{
		DBConn: db,
	}
}

func (db *model) Add(movie operation.Movie) {
	db.DBConn.Model(&operation.Movie{}).Create(&movie)
}

func (db *model) Update(movie operation.Movie) {
	db.DBConn.Model(&operation.Movie{}).Save(&movie)
}

func (db *model) Delete(movie operation.Movie) {
	db.DBConn.Model(&operation.Movie{}).Delete(&movie)
}

func (db *model) GetAll() []operation.Movie {
	var movies []operation.Movie
	db.DBConn.Model(&operation.Movie{}).Set("gorm:auto_preload", true).Find(&movies)
	return movies
}

func (db *model) GetOne(id uint64) operation.Movie {
	var movie operation.Movie
	db.DBConn.Model(&operation.Movie{}).Set("gorm:auto_preload", true).Find(&movie, id)
	return movie
}

func (db *model) AddUser(newuser user.User) {
	db.DBConn.Model(&user.User{}).Create(&newuser)
}

func (db *model) GetAllUsers() []user.User {
	var newusers []user.User
	db.DBConn.Model(&user.User{}).Find(&newusers)
	return newusers
}

func (db *model) GetUser(username string) user.User {
	var olduser user.User
	db.DBConn.Model(&user.User{}).Where("user_name = ?", username).Take(&olduser)
	return olduser
}
