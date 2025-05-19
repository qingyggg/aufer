package rpcpack

import (
	"github.com/qingyggg/aufer/biz/model/http/interact/comment"
	commentcmd "github.com/qingyggg/aufer/kitex_gen/cmd/comment"
)

func ConverComment(c *commentcmd.Comment) *comment.Comment {
	return &comment.Comment{
		Cid:           c.Cid,
		Aid:           c.Aid,
		User:          ConvertUserBase(c.User),
		Content:       c.Content,
		CreateDate:    c.CreateDate,
		FavoriteCount: c.FavoriteCount,
		ChildNum:      c.ChildNum,
		RepliedUid:    c.RepliedUid,
		IsFavorite:    c.IsFavorite,
	}
}

func ConverComments(cs []*commentcmd.Comment) (cmts []*comment.Comment) {
	for _, c := range cs {
		cmts = append(cmts, ConverComment(c))
	}
	return
}
