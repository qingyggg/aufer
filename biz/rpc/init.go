package rpc

import (
	"github.com/qingyggg/aufer/biz/model/cmd/article/articlehandler"
	"github.com/qingyggg/aufer/biz/model/cmd/comment/commenthandler"
	"github.com/qingyggg/aufer/biz/model/cmd/relation/relationhandler"
	"github.com/qingyggg/aufer/biz/model/cmd/user/userhandler"
	article_rpc "github.com/qingyggg/aufer/cmd/article/rpc"
	cmt_rpc "github.com/qingyggg/aufer/cmd/comment/rpc"
	relation_rpc "github.com/qingyggg/aufer/cmd/relation/rpc"
	user_rpc "github.com/qingyggg/aufer/cmd/user/rpc"
)

type RpcClients struct {
	RelationClient relationhandler.Client
	UserClient     userhandler.Client
	ArticleClient  articlehandler.Client
	CommentClient  commenthandler.Client
}

var Clients *RpcClients

func InitRpc() {
	c := new(RpcClients)
	c.RelationClient = relation_rpc.InitRelationRpc()
	c.UserClient = user_rpc.InitUserRpc()
	c.ArticleClient = article_rpc.InitArticleRpc()
	c.CommentClient = cmt_rpc.InitCommentRpc()
	Clients = c
}
