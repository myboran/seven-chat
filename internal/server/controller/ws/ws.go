package ws

import (
	"github.com/gorilla/websocket"
	"github.com/myboran/seven-chat/internal/server/service/chat"
	"net/http"
)

func Upgrade(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	client := chat.NewClient(r.RemoteAddr, conn)
	chat.GetManager().RegisterClient(client)
}
