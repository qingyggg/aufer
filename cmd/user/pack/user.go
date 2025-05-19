package pack

import (
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	"github.com/qingyggg/aufer/kitex_gen/cmd/user"
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
		CreatedDate:     u.CreatedAt.Format("2006-01-02"),
	}
}
