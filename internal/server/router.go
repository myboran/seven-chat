package server

import (
	"github.com/gin-gonic/gin"
	"github.com/myboran/seven-chat/internal/server/controller/ws"
	"github.com/myboran/seven-chat/internal/server/service/chat"
)

func (a *App) initRouter() {

}

func (a *App) initWS() {
	manager := chat.GetManager()
	go manager.Start()
	a.r.GET("/ws", func(c *gin.Context) {
		ws.Upgrade(c.Writer, c.Request)
	})
}
