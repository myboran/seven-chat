package chat

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/myboran/seven-chat/pkg/api"
	"log"
)

type Client struct {
	addr string
	conn *websocket.Conn
	send chan []byte
	uuid string
}

func NewClient(addr string, conn *websocket.Conn) *Client {
	client := &Client{
		addr: addr,
		conn: conn,
		send: make(chan []byte),
		uuid: uuid.New().String(),
	}
	go client.Read()
	go client.Write()
	return client
}

func (c *Client) Destroy() {
	log.Println("destroy client:", c.uuid)
	close(c.send)
	_ = c.conn.Close()
}

func (c *Client) Read() {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("ReadMessage error:", err)
			GetManager().UnregisterClient(c)
			c.Destroy()
			return
		}
		c.Process(msg)
	}
}

func (c *Client) Write() {
	var err error
	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				return
			}
			if err = c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("WriteMessage error:", err)
			}
		}
	}
}

func (c *Client) Send(msg []byte) {
	c.send <- msg
}

func (c *Client) Process(msg []byte) {
	var req api.WSRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		log.Println("process Unmarshal err: ", err, "msg: ", string(msg))
	}
	switch req.Type {
	case api.TypeSendMsg:
		msg = []byte(fmt.Sprintf("%s %s: %s", req.Data.Time, req.Data.Uuid, req.Data.Msg))
		GetManager().Broadcast(msg)
	default:
		c.Send([]byte("非法请求"))
	}
}
