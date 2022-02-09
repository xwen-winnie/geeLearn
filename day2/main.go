package main

import (
	"fmt"
	"log"
	"net/http"

)
func main(){
	enginee :=new(Enginee)//创建engine实例
	log.Fatal(http.ListenAndServe(":9999",enginee))//第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。第二个参数，则是我们基于net/http标准库实现Web框架的入口s
}


type Enginee struct {

}

func (*Enginee)ServeHTTP( w http.ResponseWriter, req *http.Request){//第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应
	switch req.URL.Path {
	case "/":
		for k,v :=range req.Header{
			fmt.Fprintf(w,"Header[%q]=%q\n",k,v)
		}
	case "/index":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)

	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s",req.URL)
	}
}