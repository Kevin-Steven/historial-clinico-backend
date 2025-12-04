package repository

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"historial-clinico-backend/signos_vitales/internal/config"
)

func NewMySQLDB(cfg config.Config) (*gorm.DB, error) {
	dsn := cfg.DBDSN
	if dsn == "" {
		log.Println("advertencia: SIGNOS_DB_DSN no está definido")
	}
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
