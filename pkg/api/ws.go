package api

import "encoding/json"

const (
	TypeSend = "send"
)

type Body struct {
	Type string `json:"type"`
	Data struct {
		Uuid string `json:"uuid"`
		Msg  string `json:"msg"`
		Time string `json:"time"`
	} `json:"data"`
}

func (b *Body) Marshal() []byte {
	bs, _ := json.Marshal(b)
	return bs
}
