package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"websocket/write2Log"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

// 使用命令
func init() {
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for WebSocket connection")
	flag.Parse()
	var err error
	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}
}

// 执行Websocket
func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))

		// 这里尝试添加日志
		fmt.Println(string(msg))
		write2Log.WriteLog("test.log", string(msg))

	}
}

// 执行js模板
func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveWS)
	r.HandleFunc("/k.js", serveFile)
	log.Fatal(http.ListenAndServe(":8080", r))
}

