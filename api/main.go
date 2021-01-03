package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shien/weightrec-backend/pkg/auth"
	"github.com/shien/weightrec-backend/pkg/user"
)

func main() {

	ur := user.Init()
	defer ur.Close()
	if err := ur.AddUser("testuser"); err != nil {
		log.Println(err)
	}

	url, err := auth.GetLoginURL()
	if err != nil {
		log.Println(err)
	}
	log.Println(url)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.GET("/user/:id/*action", func(c *gin.Context) {

		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	r.GET("/weight/:userid/data", func(c *gin.Context) {
		today := time.Now()
		year, month, day := today.Date()
		date := fmt.Sprintf(
			"%d-%d-%d",
			year, month, day,
		)
		s := c.DefaultQuery("startDate", date)
		e := c.Query("endDate")
		c.String(http.StatusOK, s, e)
	})
	r.Run()
}
