package util

import (
	"encoding/json"
	"log"
)

type MsgFromServer struct {
	Comments []string `json:"comments"`
	Reaction int      `json:"reaction"`
}

func JsonToString(jsonMsg MsgFromServer) string {
	bytes, err := json.Marshal(jsonMsg)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
