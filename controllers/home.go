package controllers

import (
	"blog-api/service/client"
	"blog-api/utils"
	log "code.google.com/log4go"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Get() {
	pageInput := c.Input().Get("page")
	limitInput := c.Input().Get("limit")
	var pageNum, limit int
	var err error
	if len(pageInput) != 0 {
		pageNum,err = strconv.Atoi(pageInput)
		if err != nil {
			pageNum = 1
			log.Error("参数类型转换错误： "+err.Error())
		}
	}
	limit = 6
	if len(limitInput) != 0 {
		limit, err = strconv.Atoi(limitInput)
		if err != nil {
			log.Error("参数类型转换错误： "+err.Error())
		}
	}

	fmt.Print("page, limit",pageNum, limit)
	topics, total, err1 := client.GetTopicsByOffsetAndLimit(pageNum, limit)

	if err1 != nil {
		log.Error("查询失败, %v", err1)
		c.Redirect("/error", 302)
	}
	//for i:=0; i<len(*topics); i++ {
	//	(*topics)[i].Content = utils.VerifyHtml((*topics)[i].Content)
	//	(*topics)[i].Content = (*topics)[i].Content[:1000]
	//}
	page := *utils.PageUtil(total, pageNum, limit, *topics)
	c.Data["Page"] = page
	c.Data["IsHome"] = true
	c.Data["PageName"] = "我的首页"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "home.html"
}
