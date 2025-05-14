package db

import (
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	mydal "github.com/qingyggg/aufer/cmd/relation/dal"
	"github.com/qingyggg/aufer/pkg/utils"
)

func AddNewFollow(follow *orm_gen.Follow) error {
	var f = mydal.Qdb.Follow
	err := f.Create(follow)
	return err
}

// DeleteFollow delete follow relation in db and update redis
func DeleteFollow(follow *orm_gen.Follow) error {
	var f = mydal.Qdb.Follow
	_, err := f.Where(f.UserID.Eq(follow.UserID), f.FollowerID.Eq(follow.FollowerID)).Delete()
	return err
}

// QueryFollowExist check the relation of user and follower
func QueryFollowExist(uHashId, fUHashId string) (bool, error) {
	var f = mydal.Qdb.Follow
	uid := utils.ConvertStringHashToByte(uHashId)
	fid := utils.ConvertStringHashToByte(fUHashId)
	count, err := f.Where(f.UserID.Eq(uid), f.FollowerID.Eq(fid)).Count()
	if err != nil {
		return false, err
	} else if count != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

// GetFollowCount query the number of users following
func GetFollowCount(uHashId string) (int64, error) {
	var f = mydal.Qdb.Follow
	count, err := f.Where(f.FollowerID.Eq(utils.ConvertStringHashToByte(uHashId))).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetFollowerCount query the number of followers of a user
func GetFollowerCount(uHashId string) (int64, error) {
	var f = mydal.Qdb.Follow
	count, err := f.Where(f.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetFollowIdList  find user_id follow id list in db
func GetFollowIdList(uHashId string) ([]string, error) {
	var f = mydal.Qdb.Follow
	follows, err := f.Where(f.FollowerID.Eq(utils.ConvertStringHashToByte(uHashId))).Find()
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range follows {
		result = append(result, utils.ConvertByteHashToString(v.UserID))
	}
	return result, nil
}

// GetFollowerIdList getFollowerIdList get follower id list in db
func GetFollowerIdList(uHashId string) ([]string, error) {
	var f = mydal.Qdb.Follow
	follows, err := f.Where(f.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Find()
	if err != nil {
		return nil, err
	}
	var result []string
	for _, v := range follows {
		result = append(result, utils.ConvertByteHashToString(v.FollowerID))
	}
	return result, nil
}
