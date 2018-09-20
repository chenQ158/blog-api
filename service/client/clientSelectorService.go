package client

import (
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/clientselector"
	"time"
)

var Selector *clientselector.EtcdClientSelector
var LoginSelector *clientselector.EtcdClientSelector
var LogoutSelector *clientselector.EtcdClientSelector
var CategorySelector *clientselector.EtcdClientSelector
var CommentSelector *clientselector.EtcdClientSelector
var TopicSelector *clientselector.EtcdClientSelector

var LoginN = "rpc.Login"
var LogoutN = "rpc.Logout"
var CategoryN = "rpc.Category"
var CommentN = "rpc.Comment"
var TopicN = "rpc.Topic"

func newSelector(etcds []string, name string) *clientselector.EtcdClientSelector {
	return clientselector.NewEtcdClientSelector(etcds, "/rpcx/"+LoginN, time.Minute, rpcx.WeightedRoundRobin, time.Minute)
}

func InitSelector(etcds []string) {
	LoginSelector = newSelector(etcds, LoginN)
	LogoutSelector = newSelector(etcds, LogoutN)
	CategorySelector = newSelector(etcds, CategoryN)
	CommentSelector = newSelector(etcds, CommentN)
	TopicSelector = newSelector(etcds, TopicN)
}

