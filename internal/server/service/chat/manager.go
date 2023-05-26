package chat

import (
	"github.com/myboran/seven-chat/pkg/api"
	"github.com/myboran/seven-chat/pkg/mtime"
	"sync"
)

var (
	manager *Manager
	once    sync.Once
)

type Manager struct {
	clients    map[string]*Client //客户端
	register   chan *Client       //注册用户
	unregister chan *Client       //注销用户
	broadcast  chan []byte        //广播
}

func GetManager() *Manager {
	once.Do(func() {
		manager = &Manager{
			clients:    map[string]*Client{},
			register:   make(chan *Client),
			unregister: make(chan *Client),
			broadcast:  make(chan []byte),
		}
	})
	return manager
}

func (m *Manager) Start() {
	for {
		select {
		case client := <-m.register:
			m.EventRegister(client)
			var body api.Body
			body.Data.Time = mtime.Now()
			body.Data.Uuid = "system"
			body.Data.Msg = "welcome " + client.uuid
			m.EventBroadcast(body.Marshal())
		case client := <-m.unregister:
			m.EventUnregister(client)
		case msg := <-m.broadcast:
			m.EventBroadcast(msg)
		}
	}
}

func (m *Manager) EventRegister(client *Client) {
	m.clients[client.uuid] = client
}
func (m *Manager) EventUnregister(client *Client) {
	delete(m.clients, client.uuid)
}

func (m *Manager) EventBroadcast(msg []byte) {
	for i := range m.clients {
		go m.clients[i].Send(msg)
	}
}

func (m *Manager) Broadcast(msg []byte) {
	m.broadcast <- msg
}

func (m *Manager) RegisterClient(client *Client) {
	m.register <- client
}

func (m *Manager) UnregisterClient(client *Client) {
	m.unregister <- client
}
