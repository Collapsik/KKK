package pages

import (
	"github.com/gin-gonic/gin"
	"kkk/structs"
	"net/http"
	"strconv"
)

func GetIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"title": "Сайт-визитка", "name": `taliban`})
}

func PostIndex(c *gin.Context) {
	fb := structs.Feedback{}
	ret := gin.H{"title": "Сайт-визитка", "name": `taliban`}
	if err := c.Bind(&fb); err != nil {
		ret["err"] = "Упс, ошибка: " + err.Error()
	} else {
		if err := structs.FeedbackSession.Insert(fb); err != nil {
			ret["err"] = "Неожиданная ошибка. Зайдите к нам попозже."
		} else {
			ret["ok"] = "Спасибо за ваш отзыв!"
		}
	}
	c.HTML(200, "index.html", ret)
}

func GetAdmin(c *gin.Context) {
	fbks := []structs.Feedback{}
	structs.FeedbackSession.Find(gin.H{}).All(&fbks)
	c.HTML(200, "admin.html", gin.H{"feedbacks": fbks})
}

func GetAdminNewPage(c *gin.Context) {
	c.HTML(200, "newpage.html", nil)
}
func PostAdminNewPage(c *gin.Context) {
	pg := structs.Page{}
	ret := gin.H{}
	if err := c.Bind(&pg); err != nil {
		ret["err"] = "Упс, ошибка: " + err.Error()
	} else {
		pg.Id = structs.GenId(structs.PageSession)
		_ = structs.PageSession.Insert(&pg)
		ret["ok"] = "Страница создана!"
	}
	c.HTML(200, "newpage.html", ret)
}

func GetAdminPages(c *gin.Context) {
	pge := []structs.Page{}
	structs.PageSession.Find(gin.H{}).All(&pge)
	c.HTML(200, "pages.html", gin.H{"page": pge})
}

func GetPage(c *gin.Context) {
	pg := structs.Page{}
	structs.PageSession.Find(gin.H{"title": c.Param("kkk")}).One(&pg)
	c.HTML(200, "tempp.html", gin.H{"page": pg})
}

func GetAdminPageDel(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("kkk"))
	structs.PageSession.Remove(gin.H{"id": i})
	c.Redirect(302, "./../pages")
}

func GetAdminPageEdit(c *gin.Context) {
	pg := structs.Page{}
	i, _ := strconv.Atoi(c.Param("kkk"))
	structs.PageSession.Find(gin.H{"id": i}).One(&pg)
	c.HTML(200, "edit.html", gin.H{"page": pg})
}

func PostAdminPageEdit(c *gin.Context) {
	pg := structs.Page{}
	ret := gin.H{}
	if err := c.Bind(&pg); err != nil {
	}
	if err := structs.PageSession.Update(gin.H{"id": pg.Id}, gin.H{"id": pg.Id, "title": pg.Title, "message": pg.Message}); err != nil {
		ret["err"] = "Ошибка базы данных" + err.Error()
	} else {
		ret["ok"] = "Страница отредактирована!"
	}
	c.HTML(200, "edit.html", ret)
}

func GetLogin(c *gin.Context) {
	cookiep, _ := c.Request.Cookie("password")
	cookiel, _ := c.Request.Cookie("login")
	if cookiel != nil {
		if cookiel.Value == structs.Login && cookiep.Value == structs.Password {
			c.Redirect(302, "./../admin")
		}
	}
	http.SetCookie(c.Writer, cookiel)
	http.SetCookie(c.Writer, cookiep)
	c.HTML(200, "login.html", nil)
}

func PostLogin(c *gin.Context) {
	cookiep, _ := c.Request.Cookie("password")
	cookiel, _ := c.Request.Cookie("login")
	us := structs.User{}
	ret := gin.H{}
	if err := c.Bind(&us); err != nil {
		ret["err"] = "Упс, ошибка: " + err.Error()
	} else {
		if us.Login == structs.Login && structs.Password == structs.Hash(us.Password) {
			cookiel = &http.Cookie{
				Name:  "login",
				Value: structs.Login,
			}
			cookiep = &http.Cookie{
				Name:  "password",
				Value: structs.Password,
			}
			http.SetCookie(c.Writer, cookiel)
			http.SetCookie(c.Writer, cookiep)
			c.Redirect(302, "./../admin")
		} else {
			ret["err"] = "Упс, ошибка: "
			c.HTML(200, "login.html", ret)
		}
	}

}
