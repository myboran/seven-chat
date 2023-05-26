package api

const (
	TypeSendMsg = "sendMsg"
)

type WSRequest struct {
	Type string `json:"type"`
	Data struct {
		Uuid string `json:"uuid"`
		Msg  string `json:"msg"`
		Time string `json:"time"`
	} `json:"data"`
}
