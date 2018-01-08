package router

import (
	"queue/model"
	"encoding/json"
	"log"
)


func generateResponseBody(errno int, msg string, data interface{}) ([]byte) {
	body := &model.ResponseBody{
	}
	body.Code = errno
	body.Message = msg
	body.Data = data
	bytes, err := json.Marshal(body)
	if (err != nil) {
		log.Printf("生成数据失败 %s", err.Error())
		return []byte(`{"code":"1", "message": "生成响应body异常", "data":[]}`)
	}
	return bytes
}

func generateSuccessBody(msg string, data interface{}) []byte {
	return generateResponseBody(model.Success, msg, data)
}

func generateFailureBody(msg string) []byte {
	return generateResponseBody(model.Failure, msg, nil)
}
