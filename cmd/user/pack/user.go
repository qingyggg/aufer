package pack

import (
	"github.com/qingyggg/aufer/biz/model/cmd/user"
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/pkg/utils"
)

func UserAssign(u *orm_gen.User) *user.UserBase {
	return &user.UserBase{
		Name:            u.UserName,
		Avatar:          utils.URLconvert(u.Avatar),
		BackgroundImage: utils.URLconvert(u.BackgroundImage),
		Signature:       u.Signature,
		Uid:             utils.ConvertByteHashToString(u.HashID),
		Email:           u.Email,
	}
}
