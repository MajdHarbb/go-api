package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// struct defining $POST data
type auth_user struct {
	User_id         string `json:"user_id"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	Dob             string `json:"dob"`
	Password        string `json:"password"`
	Password_repeat string `json:"password_repeat"`
	Date_created    string `json:"date_created"`
	Date_updated    string `json:"date_updated"`
}

//insert new user into user table
func signup(c *gin.Context) {

	var newUser auth_user

	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	//connect to mysql db
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	// catch error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	currentTime := time.Now()
	insert, err := db.Query("INSERT INTO user(user_id,username, password, date_created, date_updated) VALUES (NULL,?,?,?,?)", newUser.Username, newUser.Password, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newUser)
}

func login(c *gin.Context) {

	var user auth_user

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var loggedUser auth_user

	err = db.QueryRow("SELECT * FROM user WHERE username = ? AND password=?", user.Username, user.Password).Scan(&loggedUser.User_id, &loggedUser.Username, &loggedUser.Password, &loggedUser.Date_created, &loggedUser.Date_updated)

	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, loggedUser)
}
