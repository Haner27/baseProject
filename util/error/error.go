package error

import "fmt"

type ErrCode int

const (
	UNKNOWN        ErrCode = -1
	Success        ErrCode = 0
	Failure        ErrCode = 1
	ParamsRequired ErrCode = 2
	ParamsInvalid  ErrCode = 3

	// for user
	UserNotFound       ErrCode = 200
	UserRegisterFailed ErrCode = 201
)

func (e ErrCode) String() string {
	switch e {
	case UNKNOWN: return "未知错误"
	case Success: return "成功"
	case Failure: return "失败"
	case ParamsRequired: return "参数缺失"
	case ParamsInvalid: return "参数无效"
	case UserNotFound: return "找不到用户"
	case UserRegisterFailed: return "用户注册失败"
	default:
		return "未知错误"
	}
}

func GetErrorStr(e ErrCode) string {
	return fmt.Sprintf("%v", e)
}