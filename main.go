package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var token string
var port string
var host string
var signup string

func init() {
	flag.StringVar(&port, "p", "4000", "-p 4000")
	flag.StringVar(&host, "host", "localhost", "-host localhost")
	flag.StringVar(&signup, "signup", "/signup.xhtml", "-signup /signup.xhtml")
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gZGEgU2lsdmEiLCJjcGYiOiIyMjQ2NzQ0ODkwOCIsImV4cCI6MzYwMH0.P0TxgJPXxO6mHkfGtov8ikDjpQkVNFSzS2TZ234b7o4"
}

func main() {
	log.Println("Authorization Server Mock")
	log.Println("")
	flag.Parse()
	flag.PrintDefaults()
	log.Println("")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("pages/*")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/_record", func(c *gin.Context) {
		token = c.PostForm("token")
	})

	r.GET("/auth/oauth/authorize", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"redirect_uri": c.Query("redirect_uri")})
	})

	r.POST("/auth/oauth/sigin", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://"+host+":"+port+"/auth/oauth/approval?redirect_uri="+c.PostForm("redirect_uri")+"&access_token="+token)
	})

	r.GET("/auth/oauth/approval", func(c *gin.Context) {
		c.HTML(http.StatusOK, "approval.tmpl", gin.H{
			"callBack":    template.URL(c.Query("redirect_uri")),
			"accessToken": c.Query("access_token"),
		})
	})

	r.GET(signup, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + port)
}
