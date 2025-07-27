package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
	// 	// 移除 /debug/ 前缀
	// 	handler := http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprintf(w, "原始路径: %s\n", r.URL.Path)
	// 	}))
	// 	handler.ServeHTTP(w, r)
	// })

	http.ListenAndServe(":8080", nil)
}
