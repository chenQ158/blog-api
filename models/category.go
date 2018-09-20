package models

import "time"

// 文章分裂
type Category struct {
	Id		int		`json:"ID"`						//主键
	Title	string		`json:"TITLE"`					//分类名称
	Created	time.Time	`json:"CREATED"`			//创建时间
	TopicTime time.Time	`json:"TOPIC_TIME"`		//更新时间
	TopicCount int	`json:"TOPIC_COUNT"`	//分类下文章数量
	Topics	[]Topic
}
