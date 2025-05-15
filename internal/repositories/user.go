package repositories 

import (
	"log"
	"auth-service/internal/domain/models"
	"gorm.io/gorm"
)

type repositories struct {
	DB *gorm.DB;
}

func NewRepositories (db *gorm.DB) *repositories {
	return &repositories{DB :db}
}

func (r *repositories) GetUserByEmail(email string) (*models.Users,error){

	var user models.Users 
	err := r.DB.Where("email = %v",email).First(&user).Error
	if err != nil {
		log.Println("fail to get User by email:%s",err)
		return nil,err
	}
	return &user,nil
}