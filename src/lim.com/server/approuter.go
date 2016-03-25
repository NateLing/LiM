package main

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"lim.com"
	"lim.com/router/ws"
	"net/http"
)

var (

	//静态文件响应
	serverFile = func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	}
)

func AppRouter(mux *mux.Router) {

	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lim_com.R.HTML(w, http.StatusOK, "notFound", nil, render.HTMLOptions{Layout: ""})
	})

	//favicon.ico
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "/public/favicon.ico")
	})

	//静态资源请求
	mux.HandleFunc("/public/*", func(w http.ResponseWriter, r *http.Request) {
		serverFile(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lim_com.R.HTML(w, http.StatusOK, "index", map[string]string{
			"Title": "hello",
		})
	})

	mux.HandleFunc("/static/*", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:] + ".html"
		http.ServeFile(w, r, path)
	})

	mux.HandleFunc("/im", ws.Im)
}
