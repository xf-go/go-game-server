package main

import (
	"fmt"

	"golang.org/x/net/websocket"
)

func funcW(ws *websocket.Conn) {
	data := ws.Request().URL.Query().Get("data")
	fmt.Println("data: ", data)
	conn := &NetDataConn{
		Connection: ws,
		MD5:        "",
	}
	conn.PullFromClient()
}
