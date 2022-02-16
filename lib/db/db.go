package db

import (
	"database/sql"
	"fmt"

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
	// todo "github.com/jinzhu/configor"
	dbConfig := map[string]string{
		"User":     "root",
		"Password": "qmakzo00",
		"Name":     "tinyUrlMock_go",
		"Protocol": "tcp",
		"Params":   "charset=utf8mb4,utf8&parseTime=True&timeout=5s&readTimeout=5s&writeTimeout=5s&sql_mode=''",
	}
	// 先連上mysql
	dsn := fmt.Sprintf("%v:%v@%v/%v?%v",
		dbConfig["User"],
		dbConfig["Password"],
		dbConfig["Protocol"],
		// dbConfig.Host,
		// dbConfig.Port,
		dbConfig["Name"],
		dbConfig["Params"],
	)
	DB, err = sql.Open("mysql", dsn)

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
	if err != nil {
		panic(err)
	}
	// defer DB.Close()
}
