package operation

import (
	"time"
)

type Movie struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=10" validator:"harrypotter-check" gorm:"type:varchar(10)"`
	Description string    `json:"desc" binding:"max=20" gorm:"type:varchar(20)"`
	Trailer     string    `json:"trailer" binding:"required,url" gorm:"type:varchar(50)"`
	Price       int       `json:"price" binding:"required"`
	LeadActor   Actor     `json:"actor" binding:"required" gorm:"foreignkey:ActorID"`
	ActorID     uint64    `json:"-"`
	Created     time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated     time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Actor struct {
	FirstName string `binding:"required" gorm:"type:varchar(20)" json: "first"`
	LastName  string `json: "last" binding:"required" gorm:"type:varchar(20)"`
	Age       int    `json: "age" binding:"get=1,lte=130"`
	Email     string `json: "email" validate:"required,email" gorm:"type:varchar(20)"`
}
