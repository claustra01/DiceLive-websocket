package util

import (
	"encoding/json"
	"log"
)

type MsgFromClient struct {
	StreamId    string `json:"stream_id"`
	Comment     string `json:"comment"`
	Reaction    bool   `json:"reaction"`
	IsConnected bool   `json:"is_connected"`
}

func StringToJson(rawMsg string) MsgFromClient {
	jsonBytes := []byte(rawMsg)
	var jsonMsg MsgFromClient
	if err := json.Unmarshal(jsonBytes, &jsonMsg); err != nil {
		log.Fatal(err)
	}
	return jsonMsg
}
