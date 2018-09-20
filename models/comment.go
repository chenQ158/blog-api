package models

import "time"

type Comment struct {
	// 主键
	Id			int		`json:"ID"`
	// 文章Id
	TopicId		int		`json:"TOPICID"`
	// 评论内容
	Content		string		`json:"CONTENT"`
	// 评论人
	Name		string		`json:"NAME"`
	// 父评论，如果没有则为0
	ParentId	int		`json:"PARENT_ID"`
	// 创建时间
	Created		time.Time	`json:"CREATED"`
}
