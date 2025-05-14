package rpc_util

import (
	"errors"
	"github.com/qingyggg/aufer/biz/model/cmd/common"
	"github.com/qingyggg/aufer/pkg/errno"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *common.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *common.BaseResponse {
	return &common.BaseResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg, ServiceTime: time.Now().Unix()}
}

func CheckBaseResp(base *common.BaseResponse) error {
	if base.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(base.StatusCode, base.StatusMsg)
	}
	return nil
}
