package initializer

import "github.com/aswinjithkukku/url-moulder/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Urls{})
}
