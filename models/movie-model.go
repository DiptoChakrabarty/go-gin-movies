package models

import (
	"github.com/DiptoChakrabarty/go-gin-movies/user"

	"github.com/DiptoChakrabarty/go-gin-movies/operation"
	"github.com/jinzhu/gorm"
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
}

type model struct {
	DBConn *gorm.DB
}

func NewModelDB() MovieModel {
	db, err := gorm.Open("sqlite3", "movie.db")
	if err != nil {
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

func (db *model) GetUser(username string) user.User {
	var olduser user.User
	db.DBConn.Model(&user.User{}).Where("username = ?", username).Take(&olduser)
	return olduser
}
