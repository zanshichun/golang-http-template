package models

import (
	"encoding/json"
)

type RequestBodyModel struct{
	Body string `json:"body"`
}

type ResponseBodyModel struct{

	Code int `json:"code"`
	Message string `json:"message"`
	Body string `json:"body"`

}

// GetRespBody: 获取响应body returns：respbody []byte , err error
func (ResponseBodyModel)GetRespBody(code int,message,body string) (respbody []byte , err error){

	respBodyEntry:=&ResponseBodyModel{
		Code: code,
		Message: message,
		Body: body,
	}

	respBodyByJson,err:=json.Marshal(respBodyEntry)
	if err!=nil{
		return nil,err
	}
	
	return respBodyByJson,nil
}