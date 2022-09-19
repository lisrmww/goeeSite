package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"goser/constant"
	"time"
)

var MySQL *gorm.DB
var sqlDB *sql.DB
var err error

func init() {
	const dsn string = "root:admin123@tcp(127.0.0.1:3306)/" + constant.DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "rk_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println("connect error", err.Error())
	}
	sqlDB, err = MySQL.DB()
	if err != nil {
		fmt.Println("sqlDB, err = MySQL.DB() error", err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func CloseMysqlConn() {
	err := sqlDB.Close()
	if err != nil {
		fmt.Println("CloseMysqlConn err", err.Error())
	}
}
