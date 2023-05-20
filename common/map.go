package common

import "golang.org/x/net/websocket"

//  map構造
//  {
// 	  "stream1": [socket1, socket2, ...],
// 	  "stream2": ...,
//  }
var (
	SocketList   map[string][]*websocket.Conn = map[string][]*websocket.Conn{}
	CommentList  map[string][]string          = map[string][]string{}
	ReactionList map[string]int               = map[string]int{}
)
