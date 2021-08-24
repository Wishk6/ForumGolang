package funcsRoutes

import (
	"net/http"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	http.ServeFile(res, req, "public/html/home.html")
}

func Profile(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	http.ServeFile(res, req, "public/html/profile.html")
}

func Setting(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	http.ServeFile(res, req, "public/html/setting.html")
}

func MyTopic(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	http.ServeFile(res, req, "public/html/my-topics.html")
}

func Css(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	filename := req.URL.Query().Get("filename")

	res.Header().Set("Content-Type", "text/css")
	http.ServeFile(res, req, "public/css/"+filename)
}

func Js(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	filename := req.URL.Query().Get("filename")

	res.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(res, req, "public/js/"+filename)
}

func Assets(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	filename := req.URL.Query().Get("filename")

	// res.Header().Set("Content-Type", "application/javascript")
	http.ServeFile(res, req, "public/assets/"+filename)
}
