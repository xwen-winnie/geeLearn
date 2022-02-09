package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/",helloHander)
	http.HandleFunc("/index",indexHander)
	log.Fatal(http.ListenAndServe(":9999",nil))//第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。第二个参数，则是我们基于net/http标准库实现Web框架的入口s
}
func helloHander(w http.ResponseWriter, req *http.Request){
	for k,v :=range req.Header{
		fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
	}

}
func indexHander(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}