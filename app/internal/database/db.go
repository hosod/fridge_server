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
	db.AutoMigrate(&entity.FoodGenre{})


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
	createDummyFoodGenreData()
}
func createDummyUserData() {
	user := entity.User{Name:"Yamada", Email:"yamada@mail.com", MyFridgeID:1}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Tanaka", Email:"tanaka@mail.com", MyFridgeID:2}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Suzuki", Email:"suzuki@mail.com", MyFridgeID:2}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Sato", Email:"sato@mail.com", MyFridgeID:3}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"John", Email:"john@mail.com", MyFridgeID:4}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Bob", Email:"bob@mail.com", MyFridgeID:4}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Emily", Email:"emily@mail.com", MyFridgeID:5}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Takahashi", Email:"takahashi@mail.com", MyFridgeID:6}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Nakamura", Email:"nakamura@mail.com", MyFridgeID:6}
	if err:=db.Create(&user).Error; err!=nil {
		log.Println(err)
	}
	user = entity.User{Name:"Kondo", Email:"kondo@mail.com", MyFridgeID:6}
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
	fridge = entity.Fridge{Name:"佐藤家"}
	if err:=db.Create(&fridge).Error; err!=nil {
		log.Println(err)
	}
	fridge = entity.Fridge{Name:"留学生寮のやつ"}
	if err:=db.Create(&fridge).Error; err!=nil {
		log.Println(err)
	}
	fridge = entity.Fridge{Name:"Emily's house"}
	if err:=db.Create(&fridge).Error; err!=nil {
		log.Println(err)
	}
	fridge = entity.Fridge{Name:"社用冷蔵庫"}
	if err:=db.Create(&fridge).Error; err!=nil {
		log.Println(err)
	}
}

func createDummyFoodGenreData() {
	foodGenre := entity.FoodGenre{Name:"野菜", Unit:"個"}
	if err:=db.Create(&foodGenre).Error; err!=nil {
		log.Println(err)
	}
	foodGenre = entity.FoodGenre{Name: "飲料", Unit:"ml"}
	if err:=db.Create(&foodGenre).Error; err!=nil {
		log.Println(err)
	}
	foodGenre = entity.FoodGenre{Name: "肉", Unit:"g"}
	if err:=db.Create(&foodGenre).Error; err!=nil {
		log.Println(err)
	}
	foodGenre = entity.FoodGenre{Name: "菓子", Unit:"個"}
	if err:=db.Create(&foodGenre).Error; err!=nil {
		log.Println(err)
	}
	foodGenre = entity.FoodGenre{Name: "液体調味料", Unit:"ml"}
	if err:=db.Create(&foodGenre).Error; err!=nil {
		log.Println(err)
	}
	foodGenre = entity.FoodGenre{Name: "固体調味料", Unit:"g"}
	if err:=db.Create(&foodGenre).Error; err!=nil {
		log.Println(err)
	}
}

func createUserFridgeRelation() {
	var relations = [][]int{
		{1,2},{2,3},{1,3},{3,4},{2,6},
	}
	for _,relation:=range relations {
		var user entity.User
		var fridge entity.Fridge
		user.ID = relation[0]
		fridge.ID = relation[1]
		db.Model(&user).Association("FollowFridge").Append(&fridge)
	}
}
