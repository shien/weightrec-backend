package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shien/weightrec-backend/pkg/auth"
	"github.com/shien/weightrec-backend/pkg/csvparser"
	"github.com/shien/weightrec-backend/pkg/globalcfg"
)

const (
	cookieUserInfo = "UserInfo"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/graphs", func(c *gin.Context) {
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

	router.GET("/upload", func(c *gin.Context) {
		userinfo, err := c.Cookie(cookieUserInfo)
		if err != nil {
			c.HTML(http.StatusOK, "upload.tmpl", gin.H{
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
			c.HTML(http.StatusOK, "upload.tmpl", gin.H{
				"title":       "WeightRec",
				"mailAddress": name,
				"authstatus":  authstatus,
			})
		}
	})

	router.GET("/api/login", func(c *gin.Context) {
		url := auth.GetLoginURL()
		c.Redirect(http.StatusSeeOther, url)
	})

	router.GET("/", func(c *gin.Context) {
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

	router.GET("/login", func(c *gin.Context) {
		userinfo, err := c.Cookie(cookieUserInfo)
		if err != nil || userinfo != "" {
			c.HTML(http.StatusOK, "login.tmpl", gin.H{
				"title": "WeightRec Login",
			})
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.GET("/weight/:userid/data", func(c *gin.Context) {
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

	router.POST("/api/upload", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("csv")
		if err != nil {
			log.Println("FormFile ", err)
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}

		_, err = csvparser.Parse(file)
		if err != nil {
			log.Println("Parse Error ", err)
			c.String(http.StatusBadRequest, "Bad Request")
			return
		}

	})

	router.GET("/user/:id/*action", func(c *gin.Context) {

		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.GET("/logout", func(c *gin.Context) {
		domain := globalcfg.GetDomain()
		c.SetCookie(cookieUserInfo, "", -1, "/", domain, false, true)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.GET("/api/callback", func(c *gin.Context) {

		code := c.Query("code")
		userinfo, err := auth.GetUserInfo(code)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "internalservererror.tmpl", gin.H{
				"title": "Internal Server Error",
			})
		}

		domain := globalcfg.GetDomain()
		c.SetCookie(cookieUserInfo, userinfo, 3600, "/", domain, false, true)

		c.Redirect(http.StatusSeeOther, "/")
	})

	return router
}
