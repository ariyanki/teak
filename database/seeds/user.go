package seeds

import (
	"teak/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	SeederList["User"] = User{}
}

type User struct{}

func (seeder User) Run(db *gorm.DB) error {
	
	var user = models.User{
		Name: "Teak Admin",
		Username: "teak",
		Email:    "teak@ariyanki.net",
		Password: "teak",
	}
	user.Password, _ = seeder.HashPassword(user.Password)
	db.Save(&user)
	return nil
}

func (User) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}
