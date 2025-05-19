package rpcpack

import (
	"github.com/qingyggg/aufer/biz/model/http/basic/user"
	usercmd "github.com/qingyggg/aufer/kitex_gen/cmd/user"
)

func ConvertUser(u *usercmd.User) *user.User {
	return &user.User{
		Base: ConvertUserBase(u.Base),
		Info: &user.UserInfo{
			FollowCount:    u.Info.FollowCount,
			FollowerCount:  u.Info.FollowerCount,
			IsFollow:       u.Info.IsFollow,
			TotalFavorited: u.Info.TotalFavorited,
			WorkCount:      u.Info.WorkCount,
		},
	}
}

func ConvertUserBase(u *usercmd.UserBase) *user.UserBase {
	return &user.UserBase{
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		Name:            u.Name,
		Uid:             u.Uid,
		CreatedDate:     u.CreatedDate,
		Email:           u.Email,
	}
}

func ConvertUserBases(us []*usercmd.UserBase) (bases []*user.UserBase) {
	for _, v := range us {
		bases = append(bases, ConvertUserBase(v))
	}
	return
}
