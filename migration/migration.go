package migration

import (
	"myapp/config"
	"myapp/graph/model"
)

func MigrateTable() {
	db := config.ConnectCockroach()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if (!db.Migrator().HasTable(&model.User{})) {
		err := db.Exec("CREATE SEQUENCE user_seq").Error
		if err != nil {
			panic(err)
		}
		err = db.Migrator().AutoMigrate(&model.User{})
		if err != nil {
			panic(err)
		}
	}
}
