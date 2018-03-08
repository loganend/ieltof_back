package server

import (
	"log"
)

var roomId int = 0

const (
	roomChannelBufSize = 1
	//статусы
	roomNotActive  = "roomNotActive"
	roomNew        = "roomNew"
	roomBusy       = "roomBusy"
	roomInProgress = "roomInProgress"
	roomClose      = "roomClose"
)

type Pair struct {
	Id                    int `json:"id"`
	channelForMessage     chan Message
	//channelForDescription chan ClientSendDescriptionRoomRequest
	channelForStatus      chan string
	server                *Server
	Talker                *Talker   `json:"client,omitempty"`
	//Operator              *Operator `json:"operator,omitempty"`
	Messages              []Message `json:"messages"`
	Status                string    `json:"status,omitempty"`
	Description           string    `json:"description,omitempty"`
	Title                 string    `json:"title,omitempty"`
}

// Create new room.
func NewPair(server *Server) *Pair {

	roomId++
	ch := make(chan Message, roomChannelBufSize)
	//channelForDescription := make(chan ClientSendDescriptionRoomRequest)
	channelForStatus := make(chan string)
	messages := make([]Message, 0)
	status := roomNotActive

	return &Pair{
		Id:                    roomId,
		channelForMessage:     ch,
		channelForStatus:      channelForStatus,
		server:                server,
		//channelForDescription: channelForDescription,
		Messages:              messages,
		Status:                status}
}



// Listen Write and Read request via chanel
func (r *Pair) Listen() {
	go r.listenWrite()
}

//Send message
// func (r Room) sendMessage(response ResponseMessage) {
// 	websocket.JSON.Send(r.Operator.ws, response)
// 	websocket.JSON.Send(r.Client.ws, response)
// }

// Listen write request via chanel
func (r *Pair) listenWrite() {
	log.Println("Listening write to room")
	for {
		select {

		// отправка сообщений участникам комнаты
		//case msg := <-r.channelForMessage:
		//	r.Messages = append(r.Messages, msg)
		//	messages, _ := json.Marshal(r.Messages)
		//	response := ResponseMessage{Action: actionSendMessage, Status: "OK", Room: r.Id, Code: 200, Body: messages}
		//	log.Println(response)
		//	//r.Client.ch <- response
		//	//r.Operator.ch <- response
		//
		//	//добавление описание комнате
		//case description := <-r.channelForDescription:
		//	r.Description = description.Description
		//	r.Title = description.Title
		//	r.Status = roomNew
		//	msg, _ := json.Marshal(r)
		//	response := ResponseMessage{Action: actionSendDescriptionRoom, Status: "OK", Code: 200, Body: msg}
		//	r.Client.ch <- response
		//	r.server.broadcast(response)
		//
		//	//изменения статуса комнаты
		//case msg := <-r.channelForStatus:
		//	r.Status = msg
		//	jsonstring, _ := json.Marshal(r)
		//	response := ResponseMessage{Action: actionChangeStatusRooms, Status: "OK", Code: 200, Body: jsonstring}
		//	//websocket.JSON.Send(r.Client.ws, response)
		//	r.Client.ch <- response
		//	r.server.broadcast(response)
		}
	}
}
