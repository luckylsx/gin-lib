package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-lib/app/utils/logging"
	resp "gin-lib/app/utils/response"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

/*
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}*/

// SetUp route log record middleware
func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		//开始时间
		start := time.Now()
		startTime := start.Format("2006-01-02 15:04:05")
		//处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		var responseCode int
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			response := resp.Response{}
			err := json.Unmarshal([]byte(responseBody), &response)
			if err == nil {
				responseCode = response.Code
				responseMsg = response.Message
				responseData = response.Data
			}
		}

		//结束时间
		end := time.Now()
		endTime := end.Format("2006-01-02 15:04:05")

		// 先读取 body 中的post 参数 防止 ParseForm 之后 body 参数转义
		var requestParam string

		//ParseForm 之后 再读取body 参数会转义 所以再 ParseForm 之前读取body参数
		if c.Request.Method == "POST" {
			// ioutil.ReadAll 读取 post body 参数 之后 c.Request.Body会置空 传递到下面的方法之后 下面方法无法读取 需重使用 ioutil.NopCloser 重新新写入
			requestParamBytes, _ := ioutil.ReadAll(c.Request.Body)
			requestParam = string(requestParamBytes)

			// 重新写入 Request.Body 以便下面的方法读取
			if requestParamBytes != nil {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestParamBytes))
			}
			_ = c.Request.ParseForm()
		}

		// ParseForm 之后 可使用 Request.PostForm 获取 content-type:x-www-form-urlencode 的post 参数
		if requestParam == "" {
			requestParam = c.Request.PostForm.Encode()
		}

		//日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["request_time"] = startTime
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_referer"] = c.Request.Referer()
		accessLogMap["request_post_data"] = requestParam
		accessLogMap["request_client_ip"] = c.ClientIP()

		accessLogMap["response_time"] = endTime
		accessLogMap["response_code"] = responseCode
		accessLogMap["response_msg"] = responseMsg
		accessLogMap["response_data"] = responseData

		accessLogMap["cost_time"] = fmt.Sprintf("%vms", time.Since(start).Seconds())

		// json.Marshal 默认使用 escapeHTML：true 会使字符串中的特殊字符转义
		// accessLogJSONByte, _ := json.Marshal(accessLogMap)
		// logging.Info(string(accessLogJSONByte))

		// 自定义实现 encode 使字符串中的特殊字符不转义
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(accessLogMap)

		logging.Info(bf.String())
		// accessLogJson, _ := util.JsonEncode(accessLogMap)

		/*if f, err := os.OpenFile(config.AppAccessLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
			log.Println(err)
		} else {
			f.WriteString(accessLogJson + "\n")
		}*/
	}
}
