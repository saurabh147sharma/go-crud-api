package services

import (
	"demo-api/models"
)

func CreateUser(user models.User) {

	// // hash password
	// hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// if err != nil {

	// 	return
	// }
	// // create user
	// newUser := models.User{Email: user.Email, Password: string(hash)}

	// result := initializers.DB.Create(&newUser)

	// if result.Error != nil {

	// 	return
	// }
	// return user
}
