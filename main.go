package main

import (
	"flag"
	"net/http"

	"golang.org/x/net/websocket"
)

func init() {
	flag.Set("log_dir", "./log")
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "3")
	flag.Parse()
}

func main() {
	http.Handle("/socket", websocket.Handler(funcW))
	http.ListenAndServe(":8888", nil)
}
