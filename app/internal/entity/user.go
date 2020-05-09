package entity


//Users is user info
type Users struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Register register user info to db
// func Register(c *gin.Context) {
// // production environment
// 	// engine, _ := xorm.NewEngine("mysql", "test:test@tcp(db:3306)/test")
// // dev environment
// 	engine, _ := xorm.NewEngine("mysql", "test:test@tcp(db_dev_container:3306)/test")
// 	name := c.Query("name")
// 	mail := c.Query("mail")

// 	user := Users{Name: name, Email: mail}

// 	_, err := engine.Insert(&user)
// 	if err == nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status":  "success",
// 			"message": "success regiter user",
// 		})
// 		return
// 	}
// }