package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Like struct {
	Likes_id int `json:"likes_id"`
	User_id  int `json:"user_id"`
	Post_id  int `json:"post_id"`
}

type Like_count struct {
	Count int `json:"count"`
}

func likeCreate(c *gin.Context) {

	var newLike Like

	if err := c.BindJSON(&newLike); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "bad Input")
		return
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO likes VALUES ( Null,?,? )", newLike.User_id, newLike.Post_id)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	c.IndentedJSON(http.StatusCreated, newLike)
}

func likeDelete(c *gin.Context) {

	id := c.Param("id")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	delete, err := db.Query("DELETE FROM likes WHERE likes_id= ?", id)

	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()

	c.IndentedJSON(http.StatusOK, id)
}

func likeReadByUserID(c *gin.Context) {

	id := c.Param("id")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/facebookdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	var likes Like_count

	err = db.QueryRow("SELECT Count(*) as likes FROM likes as l JOIN post as p ON l.post_id=p.post_id WHERE p.user_id=? GROUP BY p.user_id", id).Scan(&likes.Count)

	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, likes)
}
