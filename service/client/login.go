package client

import (
	"blog-api/error"
	"blog-rpc/rpc/models"
	"context"
	"fmt"
	"github.com/smallnest/rpcx"
	"log"
)

type LoginParam struct {
	Username 	string	`json:"username"`
	Password	string	`json:"password"`
}

type LoginReply struct {
	ResultCode	string
	ResultMsg	string
	Id			int
	Username	string
	Nickname	string
}

func DoLogin(username, password string) (*LoginReply, *errors.CommonError) {
	fmt.Println("username, password: ", username, password)
	var params = rpcModels.LoginParam{Username:username, Password:password}
	var reply rpcModels.LoginReply
	var result LoginReply
	client := LoginSelector.Client
	if client == nil {
		client = rpcx.NewClient(LoginSelector)
	}
	client.Call(context.Background(), LoginN+".Dologin", &params, &reply)
	if reply.ResCode != "" {
		log.Fatal("rpc端访问错误")
		return nil, &errors.CommonError{ErrCode:"1002", ErrMsg:reply.ResMsg}
	}
	if reply.ResCode == "3001" {
		log.Println("用户名或密码错误")
		return nil, &errors.CommonError{ErrCode: "1003", ErrMsg: "用户名或密码错误"}
	}

	fmt.Println("响应结果：", reply)

	result.Id = reply.Id
	result.Username = reply.Username
	result.Nickname = reply.Nickname
	result.ResultCode = "200"
	result.ResultMsg = "success"
	return &result, nil
}
