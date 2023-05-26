package client

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/myboran/seven-chat/pkg/api"
	"github.com/myboran/seven-chat/pkg/mtime"
	"net/url"
)

var (
	addr string
)

func NewClient() {
	flag.StringVar(&addr, "addr", "", "server addr: x.x.x.x:xxx")
	flag.Parse()
	if addr == "" {
		fmt.Println("please entry -help")
		return
	}
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("connect websocket server err: ", err)
		return
	}
	exit := make(chan struct{})
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("readMessage err: ", err)
				exit <- struct{}{}
			}
			var body api.Body
			err = json.Unmarshal(msg, &body)
			if err != nil {
				fmt.Println("unmarshal err: body:", err, string(msg))
				continue
			}
			process(body)
		}
	}()

	go func() {
		for {
			var msg string
			if _, err := fmt.Scanln(&msg); err != nil {
				fmt.Println("entry err: ", err)
				continue
			}
			fmt.Print("\033[1A")
			fmt.Print("\033[K")
			var body api.Body
			body.Type = api.TypeSend
			body.Data.Time = mtime.Now()
			body.Data.Msg = msg
			if err = conn.WriteMessage(websocket.TextMessage, body.Marshal()); err != nil {
				fmt.Println("write Message err: ", err)
				continue
			}
		}
	}()
	<-exit
}

func process(body api.Body) {
	fmt.Printf("%s %s: %s\n", body.Data.Time, body.Data.Uuid, body.Data.Msg)
}
