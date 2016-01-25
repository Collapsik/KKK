package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"html/template"
)

func main() {

	Login = "Admin"
	Password = Hash("secret")

	session, err := mgo.Dial("localhost") //mongodb connect
	if err != nil {
		panic(err)
	}

	NewsSession = session.DB("mydb6").C("news")
	PageSession = session.DB("mydb6").C("page")
	UserSession = session.DB("mydb6").C("user")
	FeedbackSession = session.DB("mydb6").C("feedbacks")

	r := gin.Default()
	tmpl, err := template.ParseFiles("header.html", "footer.html", "headeradmin.html", "footeradmin.html")
	r.SetHTMLTemplate(tmpl)
	r.Static("files", "./files")
	r.LoadHTMLGlob("templates/*")

	admin := r.Group("/admin")
	admin.Use(LoginMiddleware)

	r.GET("/", GetIndex)

	r.POST("/", PostIndex)

	admin.GET("/", GetAdmin)

	admin.GET("/newpage", GetAdminNewPage)

	admin.POST("/newpage", PostAdminNewPage)

	admin.GET("/pages", GetAdminPages)

	admin.GET("/newnews", GetAdminNewNews)

	admin.POST("/newnews", PostAdminNewNews)

	admin.GET("/news", GetAdminNews)

	admin.GET("/newsedit/:kkk", GetAdminNewsEdit)

	admin.POST("/newsedit/:kkk", PostAdminNewsEdit)

	r.GET("/news/:kkk", GetNews)

	r.GET("/page/:kkk", GetPage)

	admin.GET("/pagedel/:kkk", GetAdminPageDel)

	admin.GET("/pageedit/:kkk", GetAdminPageEdit)

	admin.POST("/pageedit/:kkk", PostAdminPageEdit)

	r.GET("/news", GetNewsAll)

	r.GET("/login", GetLogin)

	r.POST("/login", PostLogin)

	admin.GET("/newsdel/:kkk", GetAdminDelNews)

	r.Run(":8080")
}
