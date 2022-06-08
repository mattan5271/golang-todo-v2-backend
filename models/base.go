package models

import (
	"fmt"
	"golang-todo-v2-backend/config"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var Db *gorm.DB

func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)

	// configから読み込んだ情報を元にデータベースに接続
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}

	// 接続したデータベースにtodosテーブルを作成
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&Todo{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created table..")
	}
}
