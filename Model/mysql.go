package Model

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var mysqldb *gorm.DB

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func InitDatabase() {
	conf, _ := os.Open("config/mysql.json")
	defer conf.Close() //执行完毕后关闭连接
	var config DBConfig
	jsonParser := json.NewDecoder(conf)
	if err := jsonParser.Decode(&config); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed creating database:%w", err)
		return
	}
	db.AutoMigrate(&User{})
	mysqldb = db
}
func DataBaseSessoin() *gorm.DB {
	return mysqldb.Session(&gorm.Session{PrepareStmt: true})
}
