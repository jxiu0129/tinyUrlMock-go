package db

import (
	"database/sql"
	"fmt"
	"tinyUrlMock-go/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// DB writer db instance
	DB *sql.DB
	// DBReader only reader db instance
	DBReader *sql.DB
	// DBGorm only gorm use
	DBGorm *gorm.DB
	// DBReaderGorm only reader gorm use
	DBReaderGorm *gorm.DB
)

func Init() {
	var err error
	dbConfig := config.Config.DB
	// 先連上mysql
	dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?%v",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Protocol,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Params,
	)
	DB, err = sql.Open("mysql", dsn)
	fmt.Printf("mysql running on %v:%v\n", dbConfig.Host, dbConfig.Port)

	if err != nil {
		panic(err)
	}

	if DB == nil {
		panic("DB nil")
	}

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	// 再把連上的DB連上orm
	DBGorm, err = gorm.Open("mysql", DB)
	DBGorm.LogMode(config.Config.DB.Debug)
	if err != nil {
		panic(err)
	}
	// defer DB.Close()
}
