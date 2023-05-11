package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("seven-chat")

	// 用户名
	username := widget.NewLabel(uuid.New().String())
	username.TextStyle = fyne.TextStyle{Bold: true}

	// 聊天记录
	messageList := container.NewVBox(
		widget.NewLabel("Welcome to the seven-chat!"),
	)
	messageListContainer := container.NewVScroll(messageList)
	messageListContainer.SetMinSize(fyne.NewSize(700, 450))

	// 消息输入框
	messageEntry := widget.NewEntry()
	sendButton := widget.NewButton("Send", func() {
		message := messageEntry.Text
		if message != "" {
			messageList.Add(widget.NewLabel("You: \n" + message))
			messageEntry.SetText("")
			messageListContainer.ScrollToBottom()
		}
	})

	// 布局
	userInfo := container.NewHBox(
		username,
	)
	inputBar := container.NewBorder(nil, nil, nil, sendButton, messageEntry)
	content := container.NewVBox(
		userInfo,
		messageListContainer,
		inputBar,
	)

	myWindow.SetContent(content)
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(700, 500))
	myWindow.ShowAndRun()
}
