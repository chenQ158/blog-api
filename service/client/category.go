package client

import (
	"blog-api/error"
	"blog-api/models"
	"blog-rpc/rpc/models"
	log "code.google.com/log4go"
	"context"
	"github.com/smallnest/rpcx"
)

func AddCat(catName string) *errors.CommonError {
	var req = rpcModels.AddCatParams{catName}
	var reply rpcModels.AddCatReply
	client := CategorySelector.Client
	if client == nil {
		client = rpcx.NewClient(CategorySelector)
	}
	client.Call(context.Background(), CategoryN+".AddCat", &req, &reply)
	if reply.ResCode != "" {
		log.Error("添加分类时，rpc端调用失败，返回结果:%v", reply)
		return &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	return nil
}

func DelCat(catId int) *errors.CommonError {
	var req = rpcModels.DelCatParams{catId}
	var reply rpcModels.DelCatReply
	client := CategorySelector.Client
	if client == nil {
		client = rpcx.NewClient(CategorySelector)
	}
	client.Call(context.Background(), CategoryN+".DelCat", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取分类列表时，rpc端调用失败，返回结果:%v", reply)
		return &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	return nil
}

func GetTopicsByCatId(id, pageNum, pageSize int) (*[]models.Topic, int, *errors.CommonError) {
	var req = rpcModels.GetTopicsByCatParams{id,pageNum, pageSize}
	var reply rpcModels.GetSummarysReply
	var result []models.Topic

	client := CategorySelector.Client
	if client == nil {
		client = rpcx.NewClient(CategorySelector)
	}
	client.Call(context.Background(), TopicN+".GetTopicsByCatId", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取分类列表时，rpc端调用失败，返回结果:%v", reply)
		return &result, 0, &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	for _,info := range reply.List {
		var topic models.Topic
		topic.Id = info.Id
		topic.Created = info.Created
		topic.Title = info.Title
		topic.Author = info.Author
		topic.Content = info.Content
		topic.Attachmemt = info.Attachment
		result = append(result, topic)
	}
	total := reply.Total
	return &result, total, nil
}

func GetAllCats() (*[]models.Category, *errors.CommonError) {
	var req = 1
	var reply rpcModels.GetAllCats
	var result []models.Category
	client := CategorySelector.Client
	if client == nil {
		client = rpcx.NewClient(CategorySelector)
	}
	client.Call(context.Background(), CategoryN+".GetAllCats", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取分类列表时，rpc端调用失败，返回结果:%v", reply)
		return &result, &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	for _,info := range reply.List {
		var cat models.Category
		cat.Id = info.Id
		cat.Title = info.CatName
		result = append(result, cat)
	}
	return &result, nil
}

func GetCats(pageNum, pageSize int) (*[]models.Category, int, *errors.CommonError) {
	var req = rpcModels.GetCatsParams{pageNum, pageSize}
	var reply *rpcModels.GetCatsReply
	var result []models.Category
	client := CategorySelector.Client
	if client == nil {
		client = rpcx.NewClient(CategorySelector)
	}
	client.Call(context.Background(), CategoryN+".GetCats", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取分类列表时，rpc端调用失败，返回结果:%v", reply)
		return &result, 0, &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	for _,cat := range reply.List {
		var c models.Category
		c.Id = cat.Id
		c.Title = cat.CatName
		c.TopicTime = cat.Updated
		c.TopicCount = cat.TopicCount
		result = append(result, c)
	}
	total := reply.Total
	return &result, total, nil
}