package main

import (
	"encoding/json"
)

type Message struct {
	//External ID
	Id int
	//Path to Image
	ImgPath string
}

func parseMessage(msg []byte) (decodedMsg *Message, err error) {
	err = json.Unmarshal(msg, &decodedMsg)

	return
}
