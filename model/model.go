package model

import (
	"baseProject/datasource"
)

func MysqlMigrate(db *datasource.MysqlDB) {
	err := db.Cli.AutoMigrate(
		&UserModel{},
	)
	if err != nil {
		panic(err)
	}
}