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
	// reactionlist := make(map[string]int)
	// commentlist := make(map[string][]string)

	return func(c echo.Context) error {

		log.Println("Serving...")
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			// 初回のメッセージを送信
			err := websocket.Message.Send(ws, "Server: Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

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
				}

				// Client からのメッセージを元に返すメッセージを作成し送信する
				err = websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", rawMsg))
				if err != nil {
					c.Logger().Error(err)
				}
			}
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
