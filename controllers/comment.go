package controllers

import (
	"blog-api/models"
	"blog-api/service/client"
	log "code.google.com/log4go"
	"fmt"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

// @router /comment/add [post]
func (c *CommentController) Post() {
	var commentInfo models.CommentInfo
	err := c.ParseForm(&commentInfo)
	var comment models.CommentInfo
	if err != nil {
		log.Error("评论：请求参数转换错误："+err.Error())
		c.Data["Error"] = "添加评论失败"
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.Data["PageName"] = "错误处理"
		c.TplName = "error.html"
		return
	} else {
		comment.Content = commentInfo.Content
		comment.Name = commentInfo.Name
		comment.TopicId = commentInfo.TopicId
		comment.ParentId = commentInfo.ParentId
		err := client.AddComment(&comment)
		if err != nil {
			log.Error("评论：添加评论失败：", err)
			c.Data["Error"] = "添加评论失败"
			c.Data["IsLogin"] = checkAccount(c.Ctx)
			c.Data["PageName"] = "错误处理"
			c.TplName = "error.html"
			return
		}

	}
	c.Redirect(fmt.Sprintf("/user/topic/show?id=%d",commentInfo.TopicId), 302)
}
