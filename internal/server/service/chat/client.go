package chat

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/myboran/seven-chat/pkg/api"
	"github.com/myboran/seven-chat/pkg/mtime"
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
	var body api.Body
	body.Data.Time = mtime.Now()
	body.Data.Uuid = "system"
	body.Data.Msg = client.uuid + " login"
	client.Send(body.Marshal())
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
	var req api.Body
	err := json.Unmarshal(msg, &req)
	if err != nil {
		log.Println("process Unmarshal err: ", err, "msg: ", string(msg))
	}
	switch req.Type {
	case api.TypeSend:
		req.Data.Uuid = c.uuid
		GetManager().Broadcast(req.Marshal())
	default:
		req.Data.Uuid = "system"
		req.Data.Msg = "非法参数"
		c.Send(req.Marshal())
	}
}
