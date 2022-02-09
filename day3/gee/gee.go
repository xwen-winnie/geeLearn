package gee

/*
首先定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法。
我们在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello，
这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。
当用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中，(*Engine).Run()方法，是 ListenAndServe 的包装。
Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，
就返回 404 NOT FOUND 。
*/
import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter,r *http.Request)

type Enginee struct{
 router  map[string] HandlerFunc
}
// New is the constructor of gee.Engine
func New() *Enginee{
	return &Enginee{router: make(map[string] HandlerFunc)}
}
func (enginee *Enginee) addRoute(method string,pattern string,handler HandlerFunc){
	key := method+"-"+pattern
	enginee.router[key]= handler
}
func (enginee *Enginee)Get(pattern string,handler HandlerFunc){
	enginee.addRoute("Get",pattern,handler)

}
func (enginee *Enginee)Post(pattern string,handler HandlerFunc){
	enginee.addRoute("Post",pattern,handler)

}
//Run defines the method to start a http server
func (enginee *Enginee)Run(addr string) error{
	return http.ListenAndServe(addr,enginee)
}
//url入口
func (enginee *Enginee)ServeHTTP( w http.ResponseWriter, req *http.Request){//第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应
	key := req.Method+"-"+req.URL.Path
	if hander,ok :=enginee.router[key];ok{
		hander(w,req)
	}else{
		fmt.Fprintf(w, "404 NOT FOUND:%s",req.URL)
	}
}