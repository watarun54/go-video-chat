package controllers

import (
	"github.com/gorilla/websocket"
	"log"
)

type WebsocketController struct{}

type WsMessage struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var (
	upgrader  = websocket.Upgrader{}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan WsMessage)
)

func NewWebsocketController() *WebsocketController {
	return &WebsocketController{}
}

func (controller *WebsocketController) HandleConnections(c Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	// クライアントを新しく登録
	clients[ws] = true

	for {
		var msg WsMessage
		// 新しいメッセージをJSONとして読み込みMessageオブジェクトにマッピングする
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 新しく受信されたメッセージをブロードキャストチャネルに送る
		broadcast <- msg
	}
	return nil
}

func (controller *WebsocketController) HandleMessages() {
	for {
		// ブロードキャストチャネルから次のメッセージを受け取る
		msg := <-broadcast
		// 現在接続しているクライアント全てにメッセージを送信する
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
	return
}
