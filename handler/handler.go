package handler

import (
	"github.com/luizengdev/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

// InitializeHandler initializes the handler and its dependencies globally
func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
