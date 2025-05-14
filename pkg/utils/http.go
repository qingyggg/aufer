package utils

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/qingyggg/aufer/biz/dal/redis"
	"github.com/qingyggg/aufer/biz/model/cmd/common"
	"github.com/qingyggg/aufer/biz/rpc"
	user_rpc "github.com/qingyggg/aufer/cmd/user/rpc"
	"github.com/qingyggg/aufer/pkg/errno"
)

// BuildBaseResp convert error and build BaseResp
func BuildBaseResp(err error) *common.BaseResponse {
	if err == nil {
		return BaseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return BaseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return BaseResp(s)
}

// BaseResp build BaseResp from error
func BaseResp(err errno.ErrNo) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
func ErrResp(c *app.RequestContext, err error) {
	resp := BuildBaseResp(err)
	c.JSON(consts.StatusInternalServerError, &common.BaseResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
	})
}

var fatalMsg = "获取Uid发生错误"

// GetUid get uid from jwt
func GetUid(c *app.RequestContext, ctx context.Context) string {
	puid, exist := c.Get("current_user_id")
	rdbUser := redis.Mrs.RdbUser
	if exist {
		//1.查找redis，找不到查找数据库，并且缓存进redis中
		err, ex := rdbUser.UHashIdExist(puid.(int64))
		if err != nil {
			hlog.Fatal(fatalMsg, err)
		}
		if ex {
			return rdbUser.UHashIdGet(puid.(int64))
		} else {
			user, err := user_rpc.QueryUserBaseByPuid(rpc.Clients.UserClient, ctx, puid.(int64))
			if err != nil {
				hlog.Fatal(fatalMsg, err)
			}
			go func() {
				err := rdbUser.UHashIdSet(puid.(int64), user.Uid)
				if err != nil {
					hlog.Warn("redis设置uhashId发生错误", err)
				}
			}()
			return user.Uid
		}
	}
	return "guest"
}
