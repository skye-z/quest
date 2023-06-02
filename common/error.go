package common

import "fmt"

type CustomError struct {
	code    int
	message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("message: %s, code: %d", e.message, e.code)
}

type CustomErrors struct {
	// 未登录
	NotLoginError CustomError
	// 令牌无效
	TokenInvalidError CustomError
	// 令牌不合法
	TokenIllegalError CustomError
	// 令牌不合法
	TokenNotAvailableError CustomError
	// 意料之外
	UnexpectedError CustomError
}

var Errors = CustomErrors{
	NotLoginError:          CustomError{10100, "请先登录"},
	TokenInvalidError:      CustomError{10100, "访问令牌无效"},
	TokenIllegalError:      CustomError{10100, "访问令牌不合法"},
	TokenNotAvailableError: CustomError{10100, "令牌无法解析"},
	UnexpectedError:        CustomError{99999, "发生意料之外的错误"},
}
