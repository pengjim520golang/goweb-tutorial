http标准库实现服务器的方式
============

> ***作者:彭劲  版权所有,转自请注明来源***

# 使用http标准库实现服务器
在go web开发当中实现网络服务器非常简单只需要使用http标准库即可,下面介绍的几种实现方法您只需要掌握其中1种即可。

在golang中实现网络需要导入`net/http`包,该包中的如下接口完成用户的请求和响应工作:

* `http.ResponseWriter` 这是一个接口,主要用于完成对用户的响应工作
* `http.Request` 这是一个结构体主要用于获取用户的相关请求

## 实现方式1:通过HandleFunc自动实现处理器

当用户请求`http://localhost:8080/blc`这个URL时`:8080`是请求的网络端口,`/blc`是请求URL的一个路由地址,不同的路由地址可以负责不同的响应逻辑

```go
package main
import (
	"fmt"
	"net/http"
)
//定义路由
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello BlockChainer",r.URL.Path)
}
func main(){
	//处理请求路由
	http.HandleFunc("/blc",handler)
	//启动服务器请求端口
	http.ListenAndServe(":8080",nil)
}
```

`http.HandleFunc("/blc",handler)` 接收到路由地址`/blc`后进行逻辑处理,`handler`其参数类型为`func (w http.ResponseWriter,r *http.Request)`是处理路由逻辑的函数签名,在golang内部它会被自动实现为一个ServeHTTP处理器,我们定义路由的处理逻辑如下:

```go
//定义路由
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello BlockChainer",r.URL.Path)
}
```

使用`http.ResponseWriter`进行请求后的响应`fmt.Fprintln`进行输出响应的时候需要使用到该接口,`http.Request`负责请求,通常是一个指针类型用于确保在请求时节省资源的开销


## 实现方式2:通过自定义处理器实现网络

如果你不想使用`HandleFunc`方法去自动实现一个ServeHTTP处理器,你可以自己定义一个处理器。这个处理器是一个结构体并且必须实现`Handler`接口。

`Handler`在http标准库被定义为接口如下:

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

`ServeHTTP(ResponseWriter, *Request)`是处理器中必须实现的方法,有了上述的概念我们把代码定义处理器如下所示:

```go
//自定义处理器
type HttpHandler struct{}

//实现Handler接口中ServeHTTP方法
func (httpHandler *HttpHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello BlockChainer by HttpHandler",r.URL.Path)
}
```

在`main`函数中调用自定义的处理器使用`http.Handle`方法,定义代码如下所示:

```go
func main(){
	//创建处理器
	myHandler := &HttpHandler{}
	//调用处理器处理路由
	http.Handle("/blc",myHandler)
	http.ListenAndServe(":8080",nil)
}
```

## 实现方式3:通过http.Server配置实现网络

上面两种实现方式都是通过`http.ListenAndServe`来实现的,http标准库还提供了详细配置来实现网络,我们修改一下刚才定义的处理器让它可以出来以下3个路由地址:

* `http://localhost:8080/blc` 实现打印blc-8080
* `http://localhost:8080/chain` 实现打印chain-8080
* `http://localhost:8080/block` 实现打印block-8080

修改处理器代码如下:

```go
//自定义处理器
type HttpHandler struct{}
//实现Handler接口中ServeHTTP方法
func (httpHandler *HttpHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//根据不同的路由请求做出对应逻辑
	result := r.URL.Path + "-8080"
	fmt.Fprintln(w,result)
}
```

并且希望以上请求必须在2秒内完成,否则超时。我们就需要通过设置`http.Server`结构体的相关属性来完成,代码如下:

```go
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
```