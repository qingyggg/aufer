package db

import (
	"github.com/qingyggg/aufer/biz/model/orm_gen"
	mydal "github.com/qingyggg/aufer/cmd/user/dal"
	"github.com/qingyggg/aufer/pkg/errno"
	"github.com/qingyggg/aufer/pkg/utils"
	"golang.org/x/exp/constraints"
)

//type User

// CreateUser create user info
func CreateUser(user *orm_gen.User) error {
	var u = mydal.Qdb.User
	err := u.Create(user)
	return err
}

// QueryUser query User by email
func QueryUser(email string) (*orm_gen.User, error) {
	var u = mydal.Qdb.User
	ex, err := CheckUserExistByEmail(email)
	if err != nil {
		return nil, err
	}
	if !ex {
		return nil, errno.UserIsNotExistErr
	}
	user, err := u.Where(u.Email.Eq(email)).Take()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// QueryUserById get user in the database by user id
func QueryUserById(userId int64) (*orm_gen.User, error) {
	var u = mydal.Qdb.User
	user, err := u.Where(u.ID.Eq(userId)).Take()
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return user, nil
}

func QueryUserByHashId(uHashId string) (*orm_gen.User, error) {
	var u = mydal.Qdb.User
	user, err := u.Where(u.HashID.Eq(utils.ConvertStringHashToByte(uHashId))).Take()
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		err := errno.UserIsNotExistErr
		return nil, err
	}
	return user, nil
}

// VerifyUser verify email and password in the db
func VerifyUser(email string, password string) (int64, string, error) {
	user, err := QueryUser(email)
	if err != nil {
		if err.Error() == "record not found" {
			return 0, "", errno.UserIsNotExistErr
		}
		return 0, "", err
	}
	if ok := utils.VerifyPassword(password, user.Password); !ok {
		err = errno.PasswordIsNotVerified
		return 0, "", err
	} else {
		return user.ID, utils.ConvertByteHashToString(user.HashID), nil
	}
}

// CheckUserExistById find if user exists
func CheckUserExistById(userId int64) (bool, error) {
	var u = mydal.Qdb.User
	count, err := u.Where(u.ID.Eq(userId)).Count()
	if err != nil {
		return false, err
	} else {
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func CheckUserExistByHashId(hashId string) (bool, error) {
	var u = mydal.Qdb.User
	count, err := u.Where(u.HashID.Eq(utils.ConvertStringHashToByte(hashId))).Count()
	if err != nil {
		return false, err
	} else {
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func CheckUserExistByEmail(email string) (bool, error) {
	var u = mydal.Qdb.User
	count, err := u.Where(u.Email.Eq(email)).Count()
	if err != nil {
		return false, err
	} else {
		if count == 1 {
			return true, nil
		} else {
			return false, nil
		}
	}
}

func UserPwdModify(uid int64, new_pwd string) error {
	var u = mydal.Qdb.User
	_, err := u.Where(u.ID.Eq(uid)).Update(u.Password, new_pwd)
	if err != nil {
		return err
	}
	return nil
}

func UserProfileModify(uHashId string, payload map[string]interface{}) error {
	var u = mydal.Qdb.User
	_, err := u.Where(u.HashID.Eq(utils.ConvertStringHashToByte(uHashId))).Updates(payload)
	return err //err =err or err=nil
}

// map[user id] : user payload
func QueryUserByIds(uids []int64) (map[int64]*orm_gen.User, error) {
	var u = mydal.Qdb.User
	//对uids进行去重操作
	uMaps := make(map[int64]*orm_gen.User)
	var uniqueIDs []int64
	for _, uid := range uids {
		uMaps[uid] = new(orm_gen.User)
	}
	for k := range uMaps {
		uniqueIDs = append(uniqueIDs, k)
	}
	users, err := u.Where(u.ID.In(uniqueIDs...)).Find()
	if err != nil {
		return nil, err
	}
	for _, cu := range users {
		uMaps[cu.ID] = cu
	}

	return uMaps, nil
}

func QueryUserByHashIds(uids []string) (uMaps map[string]*orm_gen.User, err error) {
	var u = mydal.Qdb.User
	var uidsByte [][]byte
	uMaps = make(map[string]*orm_gen.User)
	//对uids进行去重操作
	uniqueIDs := trimIds(uids)
	for _, v := range uniqueIDs {
		uidsByte = append(uidsByte, utils.ConvertStringHashToByte(v))
	}
	users, err := u.Where(u.HashID.In(uidsByte...)).Find()
	if err != nil {
		return nil, err
	}
	for _, cu := range users {
		uMaps[utils.ConvertByteHashToString(cu.HashID)] = cu
	}

	return uMaps, nil
}

func trimIds[T constraints.Ordered](uids []T) (uniqueIDs []T) {
	//对uids进行去重操作
	uMaps := make(map[T]interface{})
	for _, uid := range uids {
		uMaps[uid] = new(orm_gen.User)
	}
	for k := range uMaps {
		uniqueIDs = append(uniqueIDs, k)
	}
	return
}

func GetWorkCount(uHashId string) (count int64, err error) {
	var a = mydal.Qdb.Article
	count, err = a.Where(a.UserID.Eq(utils.ConvertStringHashToByte(uHashId))).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
