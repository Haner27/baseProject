package datasource

import (
	"baseProject/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MysqlDB struct {
	Cli *gorm.DB
}

func NewMysqlDB(conf *config.Config) *MysqlDB {
	dsn := conf.Mysql.GetDSN()
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("[dataSource]connect mysql(%s) failed: %v", dsn, err))
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: sqlDB,
		}),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Sprintf("[dataSource]connect mysql(%s) failed: %v", dsn, err))
	}
	return &MysqlDB{db}
}
