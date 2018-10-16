package main

import (
	"time"
	"fmt"
	"net/http"
)


//自定义处理器
type HttpHandler struct{}
//实现Handler接口中ServeHTTP方法
func (httpHandler *HttpHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//根据不同的路由请求做出对应逻辑
	result := r.URL.Path + "-8080"
	fmt.Fprintln(w,result)
}

func main(){
	//创建处理器
	myHandler := &HttpHandler{}
	//不同的路由请求调用处理器
	http.Handle("/blc",myHandler)
	http.Handle("/chain",myHandler)
	http.Handle("/block",myHandler)
	//设置http.Server的网络端口和超时时间
	httpSever := http.Server{
		Addr : ":8080",
		Handler:myHandler,
		ReadTimeout : 2*time.Second,
	}
	httpSever.ListenAndServe()
}