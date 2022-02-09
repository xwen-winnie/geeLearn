package main

import (
	"fmt"
	"net/http"

	"github.com/xwen-winnie/geeLearn/day3/gee"
)

func main(){
	r :=gee.New()
	r.Get("/index",func(w http.ResponseWriter,req *http.Request){
		fmt.Fprintf(w,"index is:%q",req.URL.Path)
	})
	r.Post("/",func(w http.ResponseWriter,req *http.Request){
		for k,v :=range req.Header{
			fmt.Fprintf(w,"%q header is:%q",k,v)
		}
	})
    r.Run(":9999")
}