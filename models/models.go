package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 文章提交表单实体
type TopicForm struct {
	Id				int		`form:"id"`
	Title 			string		`form:"title"`
	CatId			int		`form:"category"`
	Content			string		`form:"content"`
	Attachment		string		`form:"attachment"`
	Summary			string		`form:"summary"`
	Author			string		`form:"author"`
}

// 评论请求表单实体
type CommentInfo struct {
	TopicId			int		`form:"topicId"`
	Content			string		`form:"content"`
	Name			string		`form:"name"`
	Email			string		`form:"email"`
	ParentId		int		`form:"parentId"`
}

// 文章详情，用于页面展示
type TopicInfo struct {
	Id			int		`json:"ID"`
	Title		string		`json:"TITLE"`		//文章标题
	Attachmemt	string		`json:"ATTACHMENT"`//标签
	Updated		time.Time	`json:"UPDATED"`	//更新时间
	ReplayCount	int		`json:"REPLAY_COUNT"`		//评论数
	CategoryId	int		`json:"CATEGORY_ID"`
	CategoryTitle  string	`json:"CATEGORY_TITLE"`
	Uid			int		`json:"UID"`
	Author		string		`json:"AUTHOR"`
}

