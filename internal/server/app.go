package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	r *gin.Engine
}

func (a *App) Run() {
	go func() {
		if err := a.r.Run("0.0.0.0:7777"); err != nil {
			log.Panic("启动失败")
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("服务销毁")
}

func NewApp() *App {
	app := &App{r: gin.Default()}
	app.initRouter()
	app.initWS()
	return app
}
