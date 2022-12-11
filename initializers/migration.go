package initializers

import "demo-api/models"

func MigrateDB() {
	DB.AutoMigrate(&models.User{})
}
