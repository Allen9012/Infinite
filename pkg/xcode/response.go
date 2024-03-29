package xcode

import (
	"github.com/Allen9012/Infinite/pkg/xcode/types"
	"net/http"
)

func ErrHandler(err error) (int, any) {
	// err转换成业务自定义的code
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
