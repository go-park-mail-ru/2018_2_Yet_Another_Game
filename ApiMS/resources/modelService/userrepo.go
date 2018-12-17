package modelService

import (
	"github.com/2018_2_YetAnotherGame/ApiMS/resources/models"

	"github.com/jinzhu/gorm"
)

func FindUserByID(db *gorm.DB, id string) models.User {
	var user models.User
	db.Where("id = ?", id).First(&user)
	return user
}
