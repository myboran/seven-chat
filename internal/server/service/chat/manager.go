package chat

import (
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
	for _, v := range m.clients {
		v.Send(msg)
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
