package controllers

import (
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"strconv"
)

type WebsocketController struct{}

type WsMessage struct {
	RoomId   int
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var (
	upgrader  = websocket.Upgrader{}
	clients   = make(map[int]map[*websocket.Conn]bool)
	broadcast = make(chan WsMessage)
)

func NewWebsocketController() *WebsocketController {
	return &WebsocketController{}
}

func (controller *WebsocketController) Show(c Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	t, err := template.ParseFiles("templates/rooms/show.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
		return err
	}
	if err := t.Execute(c.Response(), struct {
		Title string
	}{
		Title: "Room - " + strconv.Itoa(id),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
		return err
	}
	return nil
}

func (controller *WebsocketController) HandleConnections(c Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	if clients[id] != nil {
		clients[id][ws] = true
	} else {
		clients[id] = map[*websocket.Conn]bool{ws: true}
	}

	for {
		var msg WsMessage
		msg.RoomId = id
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients[id], ws)

			// 誰もいなくなったroomは消す
			if len(clients[id]) == 0 {
				delete(clients, id)
			}
			break
		}
		broadcast <- msg
	}
	return nil
}

func (controller *WebsocketController) HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients[msg.RoomId] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients[msg.RoomId], client)
			}
		}
	}
	return
}
