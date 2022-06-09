package models

import (
	"golang-todo-v2-backend/config"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lib/pq"
)

var Db *gorm.DB
var err error

func init() {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	Db, err = gorm.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	// 開発環境用
	// dbConnectInfo := fmt.Sprintf(
	// 	`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
	// 	config.Config.DbUserName,
	// 	config.Config.DbUserPassword,
	// 	config.Config.DbHost,
	// 	config.Config.DbPort,
	// 	config.Config.DbName,
	// )

	// // configから読み込んだ情報を元にデータベースに接続
	// Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	// if err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	fmt.Println("Successfully connect database..")
	// }

	// // 接続したデータベースにtodosテーブルを作成
	// Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Todo{})
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("Successfully created table..")
	// }
}
