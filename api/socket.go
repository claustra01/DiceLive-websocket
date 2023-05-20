package api

import (
	"fmt"
	"log"

	"github.com/claustra01/hackz-tsumaguro-websocket/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func SocketHandler() echo.HandlerFunc {

	//  map構造
	//  {
	// 	  "stream1": [socket1, socket2, ...],
	// 	  "stream2": ...,
	//  }
	socketlist := make(map[string][]*websocket.Conn)
	commentlist := make(map[string][]string)
	reactionlist := make(map[string]int)

	return func(c echo.Context) error {

		log.Println("Serving...")
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			for {
				// Client からのメッセージを読み込む
				rawMsg := ""
				jsonMsg := util.MsgFromClient{}
				err := websocket.Message.Receive(ws, &rawMsg)
				if err != nil {
					if err.Error() == "EOF" {
						log.Println(fmt.Errorf("read %s", err))
						break
					}
					log.Println(fmt.Errorf("read %s", err))
					c.Logger().Error(err)
				}

				// Jsonに展開
				jsonMsg = util.StringToJson(rawMsg)

				// 初めて接続された時
				if jsonMsg.Comment == "" && jsonMsg.Reaction == false && jsonMsg.IsConnected == true {
					socketlist[jsonMsg.StreamId] = append(socketlist[jsonMsg.StreamId], ws)
					for _, ws_v := range socketlist[jsonMsg.StreamId] {
						sendJson := util.MsgFromServer{
							Comments: commentlist[jsonMsg.StreamId],
							Reaction: reactionlist[jsonMsg.StreamId],
						}
						err = websocket.Message.Send(ws_v, fmt.Sprintf("%s", util.JsonToString(sendJson)))
						if err != nil {
							c.Logger().Error(err)
						}
					}
				}

				// コメントが送信された時
				if jsonMsg.Comment != "" && jsonMsg.Reaction == false && jsonMsg.IsConnected == true {
					commentlist[jsonMsg.StreamId] = append(commentlist[jsonMsg.StreamId], jsonMsg.Comment)
					for _, ws_v := range socketlist[jsonMsg.StreamId] {
						sendJson := util.MsgFromServer{
							Comments: commentlist[jsonMsg.StreamId],
							Reaction: reactionlist[jsonMsg.StreamId],
						}
						err = websocket.Message.Send(ws_v, fmt.Sprintf("%s", util.JsonToString(sendJson)))
						if err != nil {
							c.Logger().Error(err)
						}
					}
				}

				// リアクションされた時
				if jsonMsg.Comment == "" && jsonMsg.Reaction == true && jsonMsg.IsConnected == true {
					reactionlist[jsonMsg.StreamId]++
					for _, ws_v := range socketlist[jsonMsg.StreamId] {
						sendJson := util.MsgFromServer{
							Comments: commentlist[jsonMsg.StreamId],
							Reaction: reactionlist[jsonMsg.StreamId],
						}
						err = websocket.Message.Send(ws_v, fmt.Sprintf("%s", util.JsonToString(sendJson)))
						if err != nil {
							c.Logger().Error(err)
						}
					}
				}
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
