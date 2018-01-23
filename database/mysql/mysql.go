package mysql

// mysql abstraction

import (
	"fmt"

	// requires by the example
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/javinc/go-kit/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// Init database
func Init() {
	i, err := gorm.Open("mysql", getConfig())
	if err != nil {
		panic(fmt.Errorf("Fatal error mysql connection: %s", err))
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
