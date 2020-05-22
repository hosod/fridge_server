package database

import (
	// "os"
	"log"
	"fmt"

	"github.com/hosod/fridge_server/app/internal/entity"
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
	connect := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", user, pass, protocol, dbname)
	db,err = gorm.Open(dialect, connect)
	if err!=nil {
		log.Fatalln(err)
	}	

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Fridge{})


	createDummyData()
}
// TestInit is initialize local database for testing
func TestInit() {
	connect := "test:test@/test?charset=utf8&parseTime=True&loc=Local"
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

// createDummyData is create some dummy data for the class
func createDummyData() {
	createDummyUserData()
	createDummyFridgeData()
	createUserFridgeRelation()
}
func createDummyUserData() {
	user := entity.User{Name:"Yamada", Email:"yamada@mail.com"}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Tanaka", Email:"tanaka@mail.com"}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Suzuki", Email:"suzuki@mail.com"}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
}

func createDummyFridgeData() {
	fridge := entity.Fridge{Name:"山田家"}
	if err:=db.Create(&fridge).Error; err!=nil {
		log.Println(err)
	}
	fridge = entity.Fridge{Name:"寮のやつ"}
	if err:=db.Create(&fridge).Error; err!=nil {
		log.Println(err)
	}
}

func createUserFridgeRelation() {
	var relations = [][]int{
		{1,1},{2,2},{3,2},
	}
	for _,relation:=range relations {
		var user entity.User
		var fridge entity.Fridge
		user.ID = relation[0]
		fridge.ID = relation[1]
		db.Model(&user).Association("Fridge").Append(&fridge)
	}
}
