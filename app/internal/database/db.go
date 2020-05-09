package database

import(
	"log"
	"fmt"

	// "github.com/hosod/fridge_server/app/internal/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Use mysql in gorm
)

var (
	db *gorm.DB
	err error
)

const (
	dialect = "mysql"
	user = "test"
	pass = "test"
	dbname = "test"
)

// Init is initialize database from main function
func Init(isdev bool) {
	var host string
	if isdev {
		host = "db_dev_container"
	} else {
		host = "db_container"
	}
	protocol := fmt.Sprintf("tcp(%s:3306)", host)
	connect := fmt.Sprintf("%s:%s@%s/%s", user, pass, protocol, dbname)
	db,err = gorm.Open(dialect, connect)
	if err!=nil {
		log.Fatalln(err)
	}	
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing database
func Close() {
	if err := db.Close(); err!=nil {
		log.Fatalln(err)
	}
}

// func autoMigration() {
// 	db.AutoMigrate(&entity.Users{})
// }