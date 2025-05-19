package service

import (
	"github.com/qingyggg/aufer/biz/rpc"
	"github.com/qingyggg/aufer/cmd/comment/dal/db"
	user_rpc "github.com/qingyggg/aufer/cmd/user/rpc"
	"github.com/qingyggg/aufer/kitex_gen/cmd/comment"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
	"github.com/qingyggg/aufer/pkg/utils"
	"sync"
)

func (c *CommentService) GetCmtList(req *comment.CardsRequest) (err error, comments []*comment.Comment) {
	if req.Degree == 1 {
		err, comments = c.GetTopCmtList(req)
	} else {
		err, comments = c.GetSubCmtList(req)
	}
	return
}

func (c *CommentService) GetTopCmtList(req *comment.CardsRequest) (err error, comments []*comment.Comment) {
	//1.请求
	firstList, err := db.GetCommentListByArticleID(c.ctx, req.Aid)
	if err != nil {
		return err, nil
	}
	err, cmts := c.getCmtList(firstList, req.MyUid)
	if err != nil {
		return err, nil
	}
	return nil, cmts
}
func (c *CommentService) GetSubCmtList(req *comment.CardsRequest) (err error, cmts []*comment.Comment) {
	firstList, err := db.GetCommentListByTopCommentID(c.ctx, req.Aid, req.Cid)
	if err != nil {
		return err, nil
	}
	err, cmts = c.getCmtList(firstList, req.MyUid)
	if err != nil {
		return err, nil
	}
	return nil, cmts
}

func (c *CommentService) getCmtList(cmtList []*db.CommentItem, myUid string) (err error, cmts []*comment.Comment) {
	//获取uid,cid
	var uids []string
	var cids []string
	var wg sync.WaitGroup
	cmts = []*comment.Comment{}
	var UInfoMaps map[string]*user.UserBase
	var CInfoMaps map[string]int32
	var CCtMaps map[string]int64
	var errChan = make(chan error, 3)
	wg.Add(3)
	for _, cmt := range cmtList {
		uids = append(uids, cmt.UserID)
		cids = append(cids, cmt.HashID)
	}

	go func() {
		defer wg.Done()
		uInfoMaps, err := user_rpc.QueryUserBases(rpc.Clients.UserClient, c.ctx, uids) //请求用户信息
		if err != nil {
			errChan <- err
			return
		}
		UInfoMaps = uInfoMaps
	}()
	go func() {
		var cInfoMaps map[string]int32
		defer wg.Done()
		if myUid != "guest" {
			err, cInfoMaps = db.GetCmtFavoriteStatusMap(cids, myUid)
			if err != nil {
				errChan <- err
				return
			}
		} else {
			for _, cid := range cids {
				cInfoMaps[cid] = 0
			}
		}
		CInfoMaps = cInfoMaps
	}()
	go func() {
		defer wg.Done()
		err, cCtMaps := db.GetCmtFavoriteCtMap(cids)
		if err != nil {
			errChan <- err
			return
		}
		CCtMaps = cCtMaps
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err, nil
		}
	}
	//3.用户信息附加,形成完整的评论表
	var isFavorite bool
	for _, cmt := range cmtList {
		if CInfoMaps[cmt.HashID] == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		curPayload := comment.Comment{
			Cid:           cmt.HashID,
			Aid:           cmt.ArticleID,
			Content:       cmt.Content,
			ChildNum:      cmt.ChildNum,
			User:          UInfoMaps[cmt.UserID],
			CreateDate:    utils.ConvertBsonTimeToString(cmt.CreateTime),
			FavoriteCount: CCtMaps[cmt.HashID],
			IsFavorite:    isFavorite,
			RepliedUid:    cmt.ParentUID,
		}
		cmts = append(cmts, &curPayload)
	}
	return nil, cmts
}
