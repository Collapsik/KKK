package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"html/template"
	"kkk/news"
	"kkk/pages"
	"kkk/structs"
)

func main() {

	structs.Login = "Admin"
	structs.Password = structs.Hash("secret")

	session, err := mgo.Dial("localhost") //mongodb connect
	if err != nil {
		panic(err)
	}

	structs.NewsSession = session.DB("mydb6").C("news")
	structs.PageSession = session.DB("mydb6").C("page")
	structs.UserSession = session.DB("mydb6").C("user")
	structs.FeedbackSession = session.DB("mydb6").C("feedbacks")

	r := gin.Default()
	tmpl, err := template.ParseFiles("header.html", "footer.html", "headeradmin.html", "footeradmin.html")
	r.SetHTMLTemplate(tmpl)
	r.Static("files", "./files")
	r.LoadHTMLGlob("templates/*")

	admin := r.Group("/admin")
	admin.Use(structs.LoginMiddleware)

	r.GET("/", pages.GetIndex)

	r.POST("/", pages.PostIndex)

	admin.GET("/", pages.GetAdmin)

	admin.GET("/newpage", pages.GetAdminNewPage)

	admin.POST("/newpage", pages.PostAdminNewPage)

	admin.GET("/pages", pages.GetAdminPages)

	admin.GET("/newnews", news.GetAdminNewNews)

	admin.POST("/newnews", news.PostAdminNewNews)

	admin.GET("/news", news.GetAdminNews)

	admin.GET("/newsedit/:kkk", news.GetAdminNewsEdit)

	admin.POST("/newsedit/:kkk", news.PostAdminNewsEdit)

	r.GET("/news/:kkk", news.GetNews)

	r.GET("/page/:kkk", pages.GetPage)

	admin.GET("/pagedel/:kkk", pages.GetAdminPageDel)

	admin.GET("/pageedit/:kkk", pages.GetAdminPageEdit)

	admin.POST("/pageedit/:kkk", pages.PostAdminPageEdit)

	r.GET("/news", news.GetNewsAll)

	r.GET("/login", pages.GetLogin)

	r.POST("/login", pages.PostLogin)

	admin.GET("/newsdel/:kkk", news.GetAdminDelNews)

	r.Run(":8080")
}
