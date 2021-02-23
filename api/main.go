package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shien/weightrec-backend/pkg/auth"
	"github.com/shien/weightrec-backend/pkg/globalcfg"
	"github.com/shien/weightrec-backend/pkg/user"
)

const (
	cookieUserInfo = "UserInfo"
)

func main() {

	ur := user.Init()
	defer ur.Close()
	if err := ur.AddUser("testuser"); err != nil {
		log.Println(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/graphs", func(c *gin.Context) {
		userinfo, err := c.Cookie(cookieUserInfo)
		if err != nil {
			c.HTML(http.StatusOK, "graphs.tmpl", gin.H{
				"title":       "WeightRec",
				"mailAddress": "My Account",
				"authstatus":  false,
			})
		} else {
			// User name is mail address
			name, err := auth.GetMailAddress(userinfo)
			authstatus := false
			if err != nil {
				log.Println("Failed to get mail address:", err)
				name = "My Account"
			} else {
				authstatus = true
			}
			c.HTML(http.StatusOK, "graphs.tmpl", gin.H{
				"title":       "WeightRec",
				"mailAddress": name,
				"authstatus":  authstatus,
			})
		}
	})

	r.GET("/", func(c *gin.Context) {
		userinfo, err := c.Cookie(cookieUserInfo)
		if err != nil {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title":       "WeightRec",
				"mailAddress": "My Account",
				"authstatus":  false,
			})
		} else {
			// User name is mail address
			name, err := auth.GetMailAddress(userinfo)
			authstatus := false
			if err != nil {
				log.Println("Failed to get mail address:", err)
				name = "My Account"
			} else {
				authstatus = true
			}
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"title":       "WeightRec",
				"mailAddress": name,
				"authstatus":  authstatus,
			})
		}
	})

	r.GET("/login", func(c *gin.Context) {
		userinfo, err := c.Cookie(cookieUserInfo)
		if err != nil || userinfo != "" {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"title": "WeightRec Login",
			})
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/logout", func(c *gin.Context) {
		domain := globalcfg.GetDomain()
		c.SetCookie(cookieUserInfo, "", -1, "/", domain, false, true)
		c.Redirect(http.StatusSeeOther, "/")
	})

	r.GET("/api/login", func(c *gin.Context) {
		url := auth.GetLoginURL()
		c.Redirect(http.StatusSeeOther, url)
	})

	r.GET("/api/callback", func(c *gin.Context) {

		code := c.Query("code")
		userinfo, err := auth.GetUserInfo(code)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "503.tmpl", gin.H{
				"title": "Internal Server Error",
			})
		}

		domain := globalcfg.GetDomain()
		c.SetCookie(cookieUserInfo, userinfo, 3600, "/", domain, false, true)

		c.Redirect(http.StatusSeeOther, "/")
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
