package mysql

// mysql abstraction

import (
	"fmt"
	"time"

	// requires by the example
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/javinc/go-kit/config"
)

var (
	db *gorm.DB
)

// Init database
func Init() {
	defer func() {
		if r := recover(); r != nil {
			log.Warn("[mysql] reconnecting...")
			time.Sleep(time.Second * 5)
			Init()
		}
	}()

	i, err := gorm.Open("mysql", getConfig())
	if err != nil {
		log.Panicf("[mysql] connection error: %s", err)
	}

	// set instance
	db = i

	// disable table name's pluralization globally
	db.SingularTable(true)
}

// Db instance
func Db() *gorm.DB {
	return db
}

func getConfig() string {
	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("mysql.user"),
		config.GetString("mysql.pass"),
		config.GetString("mysql.host"),
		config.GetString("mysql.name"))
}

// Migrate schema
func Migrate(schema []interface{}) {
	db.AutoMigrate(schema...)
}
