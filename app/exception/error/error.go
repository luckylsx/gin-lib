package error

import "net/http"

// GetError get error msg by error code
func GetError(code int) string {

	// 是否有自定义的错误码 有的话使用自定义错误码
	if v, ok := errorList[code]; ok {
		return v
	}

	// 无自定义错误码 使用 go内置的网络错误码
	if errMsg := http.StatusText(code); errMsg != "" {
		return errMsg
	}

	return "An unknown error has occurred"
}

// 定义错误码
var errorList = map[int]string{
	0:   "success",
	200: "success",
	500: "服务器错误",
}
