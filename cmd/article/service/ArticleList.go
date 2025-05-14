package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/article"
	"github.com/qingyggg/aufer/biz/model/cmd/user"
	"github.com/qingyggg/aufer/biz/rpc"
	"github.com/qingyggg/aufer/cmd/article/dal/db"
	"github.com/qingyggg/aufer/cmd/article/pack"
	comment_rpc "github.com/qingyggg/aufer/cmd/comment/rpc"
	user_rpc "github.com/qingyggg/aufer/cmd/user/rpc"
	"github.com/qingyggg/aufer/pkg/utils"
	"sync"
)

func (s *ArticleService) ArticleList(req *article.CardsRequest) (cards []*article.ArticleCard, err error) {
	var wg sync.WaitGroup
	var aids []string
	var uids []string
	var UMap map[string]*user.UserBase //用户映射
	var CMap map[string]int64          //评论数映射
	var AfMap map[string]int64         //点赞数映射
	var VMap map[string]int64          //阅读数
	var CoMap map[string]int64         //收藏数量
	var errChan = make(chan error, 5)
	defer close(errChan)
	//1.请求文章信息
	aInfos, err := db.GetArticleInfos(req.AuthorId, int(req.Offset))
	if err != nil {
		return nil, err
	}
	for _, v := range aInfos {
		aids = append(aids, utils.ConvertByteHashToString(v.HashID))
		uids = append(uids, utils.ConvertByteHashToString(v.UserID))
	}
	wg.Add(5)
	//2.请求作者信息
	go func() {
		defer wg.Done()
		uMap, err := user_rpc.QueryUserBases(rpc.Clients.UserClient, s.ctx, uids)
		if err != nil {
			errChan <- err
			return
		}
		UMap = uMap
	}()
	//3.文章被评论的数量
	go func() {
		defer wg.Done()
		cMap, err := comment_rpc.CommentCounts(rpc.Clients.CommentClient, s.ctx, aids)
		if err != nil {
			errChan <- err
			return
		}
		CMap = cMap
	}()
	//4.文章被点赞的数量
	go func() {
		defer wg.Done()
		err, afMap := db.AFavoriteCtGetByAids(aids)
		if err != nil {
			errChan <- err
			return
		}
		AfMap = afMap
	}()
	//5.文章被阅览的数量
	go func() {
		defer wg.Done()
		err, vMap := db.ViewCountGets(aids)
		if err != nil {
			errChan <- err
			return
		}
		VMap = vMap
	}()
	//6.文章的收藏数量
	go func() {
		defer wg.Done()
		err, coMap := db.ACollectCtGetByAids(aids)
		if err != nil {
			errChan <- err
			return
		}
		CoMap = coMap
	}()

	wg.Wait()
	select {
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
	default:
	}
	//处理字段
	var cur *article.ArticleCard
	for _, baseInfo := range aInfos {
		cur = new(article.ArticleCard)
		cur.Info = new(article.ArticleInfo) //初始化info
		pack.BaseInfoAssign(baseInfo, cur.Info)
		cur.Info.CollectCount = CoMap[cur.Info.Aid]
		cur.Info.ViewedCount = VMap[cur.Info.Aid]
		cur.Info.CommentCount = CMap[cur.Info.Aid]
		cur.Info.LikeCount = AfMap[cur.Info.Aid]
		cur.Author = UMap[utils.ConvertByteHashToString(baseInfo.UserID)]
		cards = append(cards, cur)
	}
	return cards, nil
}
