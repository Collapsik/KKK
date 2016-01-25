package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
)

type Feedback struct {
	Name    string `form:"name" binding:"required,min=3,max=40"`
	Title   string `form:"title" binding:"required,max=150"`
	Mail    string `form:"mail" binding:"required"`
	Message string `form:"message" binding:"required"`
}

type Page struct {
	Title   string `form:"title" binding:"required,min=3,max=40"`
	Message string `form:"message" binding:"required"`
	Id      int    `form:"id"`
}

type News struct {
	Title   string `form:"title" binding:"required,min=3,max=40"`
	Message string `form:"message" binding:"required"`
	Id      int    `form:"id"`
	Short   string
}

type User struct {
	Login    string `form:"login" binding:"required,min=5, max=40"`
	Password string `form:"password" binding:"required"`
}

var NewsSession *mgo.Collection
var PageSession *mgo.Collection
var UserSession *mgo.Collection
var FeedbackSession *mgo.Collection
var Login string
var Password string

func GenId(s *mgo.Collection) int {
	id, err := s.Count()
	if err != nil {
		fmt.Println(err.Error())
	}
	var res News
	var idI int
	for i := 0; i <= id; i++ {
		err := s.Find(gin.H{"id": i}).One(&res)
		if err != nil {
			idI = i
			break
		}
	}
	return idI
}

func Hash(stri string) string {

	h := sha512.New()
	h.Write([]byte(stri))
	sha512_hash := hex.EncodeToString(h.Sum(nil))

	h2 := sha512.New()
	h2.Write([]byte(sha512_hash))
	sha512_hash2 := hex.EncodeToString(h2.Sum(nil))

	sha512_hash2 = sha512_hash2 + sha512_hash2[:15] + stri + stri[3:]

	h3 := sha512.New()
	h3.Write([]byte(sha512_hash2))
	sha512_hash3 := hex.EncodeToString(h3.Sum(nil))

	return sha512_hash3
}

func LoginMiddleware(c *gin.Context) {
	cookiep, err := c.Request.Cookie("password")
	cookiel, err := c.Request.Cookie("login")
	if err == http.ErrNoCookie {
		c.Redirect(302, "./../login")
	} else {
		if cookiel.Value != Login && cookiep.Value != Password {
			c.Redirect(302, "./../")
		}
	}
}
