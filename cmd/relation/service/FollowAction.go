package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/relation"
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/cmd/relation/dal/db"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
)

func (s *RelationService) FollowAction(req *relation.FollowRequest) error {
	var err error
	if req.Uid == req.MyUid {
		return errno.NewErrNo(errno.ParamErrCode, "用户不可以自己关注自己")
	}
	//3.follow exist ,db action
	folRelation := orm_gen.Follow{UserID: utils.ConvertStringHashToByte(req.Uid), FollowerID: utils.ConvertStringHashToByte(req.MyUid)}
	followExist, _ := db.QueryFollowExist(req.Uid, req.MyUid)
	if req.Type == FOLLOW {
		if followExist {
			return errno.FollowRelationAlreadyExistErr
		}
		err = db.AddNewFollow(&folRelation)
	} else {
		if !followExist {
			return errno.FollowRelationNotExistErr
		}
		err = db.DeleteFollow(&folRelation)
	}
	return err
}
