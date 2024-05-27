package database

import (
	"fmt"
	"src/lib/database/migrations"

	"github.com/jinzhu/gorm"
)

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(
		&migrations.User{},
	)

	if db.Error != nil {
		panic("failed to migrate database:" + db.Error.Error())
	}

	fmt.Println("Success DB migration")
}
