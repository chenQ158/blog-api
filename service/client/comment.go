package client

import (
	"blog-api/error"
	"blog-api/models"
	"blog-rpc/rpc/models"
	"context"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/log"
)

func GetCommentsByTopicId(id int) (*[]models.Comment, int, *errors.CommonError) {
	var req = rpcModels.GetCommentsParams{id}
	var reply rpcModels.GetCommentsReply
	var result []models.Comment
	client := CommentSelector.Client
	if client == nil {
		client = rpcx.NewClient(CommentSelector)
	}
	client.Call(context.Background(), CommentN+".GetCommentsByTopicId", &req, &reply)
	if reply.ResCode != "" {
		log.Error("获取文章评论时，rpc端调用失败，返回结果:%v", reply)
		return &result, 0, &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	for _,comm := range reply.List {
		var info models.Comment
		info.Id = comm.Id
		info.TopicId = comm.TopicId
		info.Content = comm.Content
		info.Name = comm.Name
		info.ParentId = comm.ParentId
		info.Created = comm.Created
		result = append(result, info)
	}
	total := reply.Total
	return &result, total, nil
}

func AddComment(comm *models.CommentInfo) *errors.CommonError {
	var req = rpcModels.CommentInfo{
		TopicId: comm.TopicId,
		ParentId: comm.ParentId,
		Content: comm.Content,
		Name: comm.Name,
		Email: comm.Email,
	}
	var reply rpcModels.AddCommentReply
	client := CommentSelector.Client
	if client == nil {
		client = rpcx.NewClient(CommentSelector)
	}
	client.Call(context.Background(), CommentN+".AddComment", &req, &reply)
	if reply.ResCode != "" {
		log.Error("添加评论时，rpc端调用失败，返回结果:%v", reply)
		return &errors.CommonError{"1001", "数据中心错误，请联系管理员"}
	}
	log.Debug("添加了评论:%v",comm)
	return nil
}
