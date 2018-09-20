package models

import "time"

// 发表文章
type Topic struct {
	Id			int		`json:"ID"`			//主键
	//Uid			int		`json:"USER_ID"`//用户id
	Title		string		`json:"TITLE"`		//文章标题
	Content		string		`json:"CONTENT"`//文章内容
	Attachmemt	string		`json:"ATTACHMENT"`//标签
	Created		time.Time	`json"CREATED"`	//创建时间
	//Views		int		`json"VIEWS"`		//浏览数
	Author		string		`json"AUTHOR"`	//作者
	//ReplayTime	time.Time	`json"REPLAY_TIME"`	//最后评论时间
	//ReplayCount	int		`json"REPLAY_COUNT"` 	//评论数
	//ReplayLastUserId int	`json"REPLAY_LAST_USER_ID"`	//最后评论用户Id
	//CategoryId	int		`json:"CATEGORY_ID"`	//所属分类
}

type TopicShow struct {
	Id            int        `json:"ID"`
	Title         string     `json:"TITLE"`
	Attachmemt    string     `json:"ATTACHMENT"`
	Updated       time.Time  `json:"UPDATED"`
	ReplayCount   int        `json:"REPLAY_COUNT"`
	CategoryId    int        `json:"CATEGORY_ID"`
	CategoryTitle string     `json:"CATEGORY_TITLE"`
	Uid           int        `json:"UID"`
	Author        string     `json:"AUTHOR"`
	Content		  string	 `json:"CONTENT"`
}
