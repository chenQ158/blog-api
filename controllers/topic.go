package controllers

import (
	"blog-api/constant"
	"blog-api/models"
	"blog-api/service/client"
	"blog-api/utils"
	log "code.google.com/log4go"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type TopicController struct {
	beego.Controller
}

// @router /user/topic/del [get]
func (this *TopicController) Del() {
	idStr := this.Input().Get("id")
	catIdStr := this.Input().Get("category")
	id,err := strconv.Atoi(idStr)
	categoryId,err := strconv.Atoi(catIdStr)
	if err != nil {
		log.Error(err)
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}
	err1 := client.DelTopic(id, categoryId)
	if err1 != nil {
		fmt.Println("删除文章失败："+err1.Error())
		this.Data["Error"] = constant.TOPIC_DEL_ERROR
		this.Redirect("/error", 302)
		return
	}

	this.Redirect("/user/topic/list", 302)
}

// @router /user/topic/add [get]
func (this *TopicController) ToAdd() {
	cats,err := client.GetAllCats()
	if err != nil {
		log.Error(err)
		this.Data["Error"] = constant.CATEGORY_LIST_ERROR
		this.Redirect("/error", 302)
		return
	}
	this.Data["Categories"] = cats
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.Data["PageName"] = "添加文章"
	this.TplName = "article_add.html"
}

// @router /user/topic/list [get]
func (this *TopicController) List() {
	pageStr := this.Input().Get("page")
	var page = 1
	var err error
	if len(pageStr) != 0 {
		page,err = strconv.Atoi(pageStr)
		if err != nil {
			log.Error(err)
			this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
			this.Redirect("/error", 302)
			return
		}

	}
	infoList,total,err1 := client.GetTopics(page, 10)
	if err1 != nil {
		log.Error(err1)
		this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
		this.Redirect("/error", 302)
		return
	}
	pageList := utils.PageUtil(total, page,10, infoList)
	this.Data["Page"] = pageList
	this.Data["InfoList"] = infoList
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.Data["PageName"] = "文章"

	this.TplName = "topic.html"
}

// @router /user/topic/edit [get]
func (this *TopicController) ToEdit() {
	idStr := this.Input().Get("id")
	if len(idStr) != 0 {
		id,err := strconv.Atoi(idStr)
		if err != nil {
			log.Error(err)
			this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
			this.Redirect("/error", 302)
			return
		}
		topic,err1 := client.GetTopicById(id)
		if err1 != nil {
			log.Error(err1)
			this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
			this.Redirect("/error", 302)
			return
		}

		//cats,err1 := client.GetAllCats()
		//if err1 != nil {
		//	log.Error(err1)
		//	this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
		//	this.Redirect("/error", 302)
		//	return
		//}
		topic.Content = utils.VerifyHtml(topic.Content)
		//this.Data["Categories"] = cats
		this.Data["Topic"] = topic
		this.Data["IsLogin"] = checkAccount(this.Ctx)
		this.Data["IsTopic"] = true
		this.Data["PageName"] = "编辑文章"
		this.TplName = "article_edit.html"
		return
	}
	// 参数错误不能编辑
	this.Data["Error"] = "文章不存在，无法修改"
	this.Redirect("/error", 302)
}

// @router /user/topic/show [get]
func (this *TopicController) Show() {
	idStr := this.Input().Get("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err)
		this.Data["Error"] = constant.VARIABLE_EXCHANGE_ERROR
		this.Redirect("/error", 302)
		return
	}
	topic, err1 := client.GetTopicDetailsById(id)
	if err1 != nil {
		log.Error(err1)
		this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
		this.Redirect("/error", 302)
		return
	}
	topic.Content = utils.VerifyHtml(topic.Content)
	comments, _, err1 := client.GetCommentsByTopicId(id)
	if err1 != nil {
		log.Error(err1)
		this.Data["Error"] = constant.VARIABLE_SEARCH_ERROR
		this.Redirect("/error", 302)
		return
	}
	this.Data["Comments"] = *comments
	this.Data["Topic"] = topic
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.Data["PageName"] = "文章详情"
	this.TplName = "article.html"
}

// @router /user/topic/search [get]
func (this *TopicController) Search() {
	keyword := this.Input().Get("keyword")
	infoList, total, err :=  client.GetTopicsByKeyword(keyword)
	if err != nil {
		log.Error("获取文章失败！",err)
		this.Data["Error"] = constant.TOPIC_LIST_ERROR
		this.Redirect("/error", 302)
		return
	}
	page := utils.PageUtil(total, 0, 10, *infoList)
	this.Data["IsSearch"] = true
	this.Data["Page"] = page
	this.Data["InfoList"] = infoList
	this.Data["PageName"] = "文章"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
}

// @router /user/topic/add [post]
func (this *TopicController) Add() {
	form := models.TopicForm{}
	if err := this.ParseForm(&form);err!=nil {
		log.Error(constant.FORM_PARSE_ERROR)
		return
	}

	topic := models.TopicForm{Title:form.Title, Attachment:form.Attachment, Content:form.Content, Author:form.Author, CatId:form.CatId, Summary:form.Summary}
	client.AddTopic(topic)
	this.Redirect("/user/topic/list", 302)
}

// @router /user/topic/update [post]
func (this *TopicController) Update() {
	form := models.TopicForm{}
	if err := this.ParseForm(&form);err!=nil {
		log.Error(constant.FORM_PARSE_ERROR)
		return
	}
	client.UpdateTopic(form)
	this.Redirect("/user/topic/list", 302)
}
