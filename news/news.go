package news

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kkk/structs"
	"strconv"
)

func GetAdminNewNews(c *gin.Context) {
	c.HTML(200, "newnews.html", nil)
}

func PostAdminNewNews(c *gin.Context) {
	nw := structs.News{}
	ret := gin.H{}
	if err := c.Bind(&nw); err != nil {
		ret["err"] = "Упс, ошибка: " + err.Error()
	} else {
		nw.Id = structs.GenId(structs.NewsSession)
		if len(nw.Message) < 300 {
			nw.Short = nw.Message
		} else {
			nw.Short = nw.Message[:300]
		}
		if err := structs.NewsSession.Insert(nw); err != nil {
			ret["err"] = "Ошибка базы данных" + err.Error()
		} else {
			ret["ok"] = "Новость создана!"
		}
	}

	c.HTML(200, "newnews.html", ret)
}

func GetAdminNews(c *gin.Context) {
	nwe := []structs.News{}
	structs.NewsSession.Find(gin.H{}).All(&nwe)
	c.HTML(200, "news.html", gin.H{"news": nwe})
}

func GetAdminDelNews(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("kkk"))
	structs.NewsSession.Remove(gin.H{"id": i})
	c.Redirect(302, "./../news")
}

func GetAdminNewsEdit(c *gin.Context) {
	nw := structs.News{}
	i, _ := strconv.Atoi(c.Param("kkk"))
	structs.NewsSession.Find(gin.H{"id": i}).One(&nw)
	fmt.Println(nw.Short, nw.Id)
	c.HTML(200, "editn.html", gin.H{"news": nw})
}

func PostAdminNewsEdit(c *gin.Context) {
	nw := structs.News{}
	ret := gin.H{}
	if err := c.Bind(&nw); err != nil {

	}
	if len(nw.Message) < 300 {
		nw.Short = nw.Message
	} else {
		nw.Short = nw.Message[:300]
	}
	if err := structs.NewsSession.Update(gin.H{"id": nw.Id}, gin.H{"id": nw.Id, "title": nw.Title, "message": nw.Message, "short": nw.Short}); err != nil {
		ret["err"] = "Ошибка базы данных" + err.Error()
	} else {
		ret["ok"] = "Новость отредактирована!"
	}
	c.HTML(200, "editn.html", ret)
}

func GetNews(c *gin.Context) {
	nw := structs.News{}
	i, _ := strconv.Atoi(c.Param("kkk"))
	structs.NewsSession.Find(gin.H{"id": i}).One(&nw)
	c.HTML(200, "temp.html", gin.H{"news": nw})
}

func GetNewsAll(c *gin.Context) {
	nw := []structs.News{}
	structs.NewsSession.Find(gin.H{}).Sort("-id").Limit(20).All(&nw)
	c.HTML(200, "news20.html", gin.H{"news": nw})
}
