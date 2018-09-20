package client

import (
	"blog-api/error"
	"blog-api/models"
	"blog-rpc/rpc/models"
	log "code.google.com/log4go"
	"context"
	"github.com/smallnest/rpcx"
)

func GetTopicsCount() int {
	return 0
}

func GetTopicsByOffsetAndLimit(pageNum, pageSize int) (*[]models.Topic, int, *errors.CommonError) {
	var req = rpcModels.GetTopicsParams{pageNum, pageSize}
	var reply rpcModels.GetSummarysReply
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".GetSummarysByOffsetAndLimit", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取文章分页时rpc端调用错误,rpc端返回信息:%v", reply)
		return nil, 0, &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	var list []models.Topic
	for _,t := range reply.List {
		var topic models.Topic
		topic.Id = t.Id
		topic.Title = t.Title
		topic.Author = t.Author
		topic.Content = t.Content
		topic.Created = t.Created
		list = append(list, topic)
	}
	total := reply.Total
	return &list, total, nil
}

func UpdateTopic(form models.TopicForm) *errors.CommonError {
	var req = rpcModels.TopicForm{
		Id: form.Id,
		Title: form.Title,
		Content: form.Content,
		Attachment: form.Attachment,
	}
	var reply rpcModels.UpdateTopicReply
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".DelTopic", &req, &reply)
	if reply.ResCode != "" {
		log.Error("更新文章时rpc端调用错误,rpc端返回信息:%v", reply)
		return &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	return nil
}

func AddTopic(topic models.TopicForm) *errors.CommonError {
	var req = rpcModels.TopicForm{
		Title: topic.Title,
		Content: topic.Content,
		CatId: topic.CatId,
		Attachment: topic.Attachment,
		Author: topic.Author,
		Summary: topic.Summary,
	}
	var reply *rpcModels.AddTopicReply
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".AddTopic", &req, &reply)
	if reply.ResCode != "" {
		log.Error("添加文章时rpc端调用错误,rpc端返回信息:%v", reply)
		return &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	return nil
}

func GetTopics(pageNum, pageSize int) (*[]models.TopicInfo, int, *errors.CommonError) {
	var req = rpcModels.GetTopicsParams{pageNum, pageSize}
	var reply rpcModels.GetTopicsReply
	var result []models.TopicInfo
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".GetTopics", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取文章分页详情信息时rpc端调用错误,rpc端返回信息:%v", reply)
		return nil, 0, &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	for _, e := range reply.List {
		var info models.TopicInfo
		info.Id = e.Id
		info.Title = e.Title
		info.Author = e.Author
		info.Attachmemt = e.Attachmemt
		info.Updated = e.Updated
		info.CategoryTitle = e.CategoryTitle
		info.CategoryId = e.CategoryId
		info.Uid = e.Uid
		info.ReplayCount = e.ReplayCount
		result = append(result, info)
	}
	total := reply.Total
	return &result, total, nil
}

func GetTopicsByKeyword(keyword string) (*[]models.TopicInfo, int, *errors.CommonError) {
	var req = rpcModels.SearTopicParams{keyword}
	var reply rpcModels.GetTopicsReply
	var result []models.TopicInfo
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".GetTopicsByKeyword", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取文章分页详情信息时rpc端调用错误,rpc端返回信息:%v", reply)
		return nil, 0, &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	for _, e := range reply.List {
		var info models.TopicInfo
		info.Id = e.Id
		info.Title = e.Title
		info.Author = e.Author
		info.Attachmemt = e.Attachmemt
		info.Updated = e.Updated
		info.CategoryTitle = e.CategoryTitle
		info.CategoryId = e.CategoryId
		info.Uid = e.Uid
		info.ReplayCount = e.ReplayCount
		result = append(result, info)
	}
	total := reply.Total
	return &result, total, nil
}

func GetTopicDetailsById(id int) (*models.TopicShow, *errors.CommonError) {
	var req = rpcModels.GetTopicParam{Id: id}
	var reply rpcModels.GetTopicDetailsReply
	var result models.TopicShow
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".GetTopicDetailsById", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取指定文章时rpc端调用错误,rpc端返回信息:%v", reply)
		return &result, &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	result.Id = reply.Id
	result.Content = reply.Content
	result.Title = reply.Title
	result.Author = reply.Author
	result.Uid = reply.Uid
	result.ReplayCount = reply.ReplayCount
	result.Updated = reply.Updated
	result.Attachmemt = reply.Attachmemt
	result.CategoryId = reply.CategoryId
	result.CategoryTitle = reply.CategoryTitle
	return &result, nil
}

func DelTopic(id, cid int) *errors.CommonError {
	var req = rpcModels.DelTopicParams{Id:id, CatId:cid}
	var reply rpcModels.DelTopicReply
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".DelTopic", &req, &reply)
	if reply.ResCode != "" {
		log.Error("删除文章时rpc端调用错误,rpc端返回信息:%v", reply)
		return &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	return nil
}

func GetTopicById(id int) (*models.Topic, *errors.CommonError) {
	var req = rpcModels.GetTopicParam{Id: id}
	var reply rpcModels.GetTopic
	var result models.Topic
	client := TopicSelector.Client
	if client == nil {
		client = rpcx.NewClient(TopicSelector)
	}
	client.Call(context.Background(), TopicN+".GetTopicById", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取指定文章时rpc端调用错误,rpc端返回信息:%v", reply)
		return &result, &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}
	result.Id = reply.Id
	result.Author = reply.Author
	result.Attachmemt = reply.Attachment
	result.Content = reply.Content
	result.Title = reply.Title
	return &result, nil
}