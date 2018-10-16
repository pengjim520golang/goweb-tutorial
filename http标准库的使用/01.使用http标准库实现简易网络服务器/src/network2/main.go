package main

import (
	"fmt"
	"net/http"
)


//自定义处理器
type HttpHandler struct{}

//实现Handler接口中ServeHTTP方法
func (httpHandler *HttpHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello BlockChainer by HttpHandler",r.URL.Path)
}

func main(){
	//创建处理器
	myHandler := &HttpHandler{}
	//调用处理器处理路由
	http.Handle("/blc",myHandler)
	http.ListenAndServe(":8080",nil)
}