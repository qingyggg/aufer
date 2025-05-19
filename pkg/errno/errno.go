package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 10000
	ParamErrCode
	AuthorizationFailedErrCode
	UserAlreadyExistErrCode
	UserIsNotExistErrCode
	FollowRelationAlreadyExistErrCode
	FollowRelationNotExistErrCode
	ArticleIsNotExistErrCode
	CommentIsNotExistErrCode
	CollectAlreadyExistErrCode
	CollectIsNotExistErrCode
	NotifyIsNotExistErrCode
)

const (
	SuccessMsg                = "Success"
	ServerErrMsg              = "Service is unable to start successfully"
	ParamErrMsg               = "Wrong Parameter has been given"
	UserIsNotExistErrMsg      = "user is not exist"
	PasswordIsNotVerifiedMsg  = "email or password not verified"
	ArticleIsNotExistErrMsg   = "article is not exist"
	CommentIsNotExistErrMsg   = "comment is not exist"
	CollectAlreadyExistErrMsg = "collect is already exist"
	CollectIsNotExistErrMsg   = "collect is not exist"
	NotifyIsNotExistErrMsg    = "notify is not exist"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                       = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr                    = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamErr                      = NewErrNo(ParamErrCode, ParamErrMsg)
	UserAlreadyExistErr           = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr        = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserIsNotExistErr             = NewErrNo(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerified         = NewErrNo(AuthorizationFailedErrCode, PasswordIsNotVerifiedMsg)
	FollowRelationAlreadyExistErr = NewErrNo(FollowRelationAlreadyExistErrCode, "Follow Relation already exist")
	FollowRelationNotExistErr     = NewErrNo(FollowRelationNotExistErrCode, "Follow Relation does not exist")
	ArticleIsNotExistErr          = NewErrNo(ArticleIsNotExistErrCode, ArticleIsNotExistErrMsg)
	CommentIsNotExistErr          = NewErrNo(CommentIsNotExistErrCode, CommentIsNotExistErrMsg)
	CollectAlreadyExistErr        = NewErrNo(CollectAlreadyExistErrCode, CollectAlreadyExistErrMsg)
	CollectIsNotExistErr          = NewErrNo(CollectIsNotExistErrCode, CollectIsNotExistErrMsg)
	NotifyIsNotExistErr           = NewErrNo(NotifyIsNotExistErrCode, NotifyIsNotExistErrMsg)
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
