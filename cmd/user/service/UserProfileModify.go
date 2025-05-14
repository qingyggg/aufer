package service

import (
	"github.com/qingyggg/aufer/biz/model/cmd/user"
	"github.com/qingyggg/aufer/cmd/user/dal/db"
	"github.com/qingyggg/aufer/pkg/utils"
)

func (s *UserService) UserProfileModify(req *user.ProfileModRequest) (error, string) {
	profile := map[string]interface{}{
		"signature":        req.Signature,
		"avatar":           utils.UrlConvertReverse(req.Avatar),
		"background_image": utils.UrlConvertReverse(req.BackgroundImage),
	}
	// 过滤空字符串的字段
	for key, value := range profile {
		if value == "" {
			delete(profile, key)
		}
	}
	err := db.UserProfileModify(req.Uid, profile)
	if err != nil {
		return err, ""
	}
	return nil, req.Uid
}
