package AuthService

import (
	"Helper"
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	AccessToken  string
	RefreshToken string
	ClientId     uint
	UserId       uint
	Status       uint
}

db.Model(&user).Related(&profile)