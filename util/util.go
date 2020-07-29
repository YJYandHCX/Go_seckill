package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//import (
//	"encoding/json"
//	"log"
//	"net/http"
//)

type ResponseData struct {
	Code int
	Msg  string
	Data interface{}
}

func Resp(writer http.ResponseWriter, code int, msg string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	rep := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(rep)
	if err != nil {
		log.Panicln(err.Error())
	}

	writer.Write(ret)
	//返回jsion Ok

}
