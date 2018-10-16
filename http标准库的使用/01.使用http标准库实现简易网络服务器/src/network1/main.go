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