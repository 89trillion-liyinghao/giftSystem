package testUtil

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
)

//post请求
func Post(uri string, router *gin.Engine) []byte {

	//构造post请求
	req := httptest.NewRequest("POST",uri,nil)

	//初始化响应
	w := httptest.NewRecorder()

	//调用相应的handler接口
	router.ServeHTTP(w,req)

	//提取响应
	result := w.Result()
	defer result.Body.Close()

	//读取body
	body,_ := ioutil.ReadAll(result.Body)
	return body
}

//post请求
func PostJson(uri string, param map[string]interface{}, router *gin.Engine) []byte {
	// 将参数转化为json比特流
	jsonByte,_ := json.Marshal(param)

	// 构造post请求，json数据以请求body的形式传递
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body,_ := ioutil.ReadAll(result.Body)
	return body
}