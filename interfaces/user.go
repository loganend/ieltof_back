package interfaces

import (
	"github.com/gorilla/websocket"
	"github.com/ieltof/server"
	"log"
	"time"
	"io"
	"encoding/json"
	"github.com/ieltof/domain"
	"fmt"
)

const channelBufSize = 100

type User struct {
	Id     uint32
	ws     *websocket.Conn
	hub    *Hub
	ch     chan server.ResponseMessage
	doneCh chan bool
}

func NewUser(hub *Hub, ws *websocket.Conn, Id uint32) *User {

	if ws == nil {
		panic("ws cannot be nil")
	}

	doneCh := make(chan bool)
	ch := make(chan server.ResponseMessage, channelBufSize)
	return &User{Id, ws, hub, ch, doneCh}
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (c *User) Listen() {

	go c.listenWrite()
	c.listenRead()
}

func (u *User) listenWrite() {
	log.Println("Listening write to user")
	//defer func() {
	//	u.hub.unregister <- u
	//	u.ws.Close()
	//}()

	//u.ws.SetReadLimit(maxMessageSize)
	//u.ws.SetReadDeadline(time.Now().Add(pongWait))
	//u.ws.SetPongHandler(func(string) error {
	//	u.ws.SetReadDeadline(time.Now().Add(pongWait));
	//	return nil
	//})

	for {
		select {
		case msg := <-u.ch:
			//log.Println("Send:", msg, c.ws)

			b, err := json.Marshal(msg)
			if err != nil {
				fmt.Println(err)
				return
			}
			u.ws.WriteMessage(websocket.TextMessage, b)
			//websocket.JSON.Send(c.ws, msg)
		}
	}
}

func (u *User) listenRead() {
	log.Println("Listening read from client")
	//ticker := time.NewTicker(pingPeriod)

	//defer func() {
	//	ticker.Stop()
	//	u.ws.Close()
	//}()

	for {
		select {
		//case message, ok := <-u.send:
		//	if !ok {
		//		u.write(websocket.CloseMessage, []byte{})
		//		return
		//	}
		//	if err := u.write(websocket.TextMessage, message); err != nil {
		//		return
		//	}

		//case <-u.doneCh:
		//	u.hub.Unregister(u)
		//	u.doneCh <- true // for listenWrite method
		//	return

		//case <-ticker.C:
		//	if err := u.write(websocket.PingMessage, []byte{}); err != nil {
		//		return
		//	}

		default:
			_, message, err := u.ws.ReadMessage()

			if err == io.EOF {
				//u.doneCh <- true
			} else if err != nil {
				//u.hub.Err(err)
			}

			var msg server.RequestMessage
			err = json.Unmarshal(message, &msg)
			//if err != nil {
			//	u.hub.Err(err)
			//}

			switch msg.Action {
			case initMessage:
				log.Println(initMessage)
				var message InitMessage
				err := json.Unmarshal(msg.Body, &message)
				if !server.CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := server.ResponseMessage{Action: initMessage, Status: "Invalid Request", Code: 403}
					u.ch <- msg
				}

				u.Id = message.UserId;
				u.hub.register <- u
				break;

			case actionSendMessage:

				log.Println(actionSendMessage)
				var message Message
				err := json.Unmarshal(msg.Body, &message)
				if !server.CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := server.ResponseMessage{Action: actionSendMessage, Status: "Invalid Request", Code: 403}
					u.ch <- msg
				}

				dbMessage := domain.Message{0, message.FromId, message.DialogId, message.Text, 0}
				dbMsg, err := InteratorInstance.NewMessage(dbMessage)
				if err != nil {
					msg := server.ResponseMessage{Action: actionSendMessage, Status: "Invalid Request", Code: 403}
					u.ch <- msg
				}

				var sign bool;
				addressee, sign := u.hub.users[message.ToId]
				if !sign {
					b, err := json.Marshal(dbMsg)
					if err != nil {
						fmt.Println(err)
						return
					}
					msg := server.ResponseMessage{Action: actionSendMessage, Status: "OK", Code: 200, Body: b}
					u.ch <- msg
					//msg := server.ResponseMessage{Action: actionSendMessage, Status: "OK", Code: 200}
					//u.ch <- msg
					break;
				}

				b, err := json.Marshal(dbMsg)
				if err != nil {
					fmt.Println(err)
					return
				}
				msg := server.ResponseMessage{Action: actionSendMessage, Status: "OK", Code: 200, Body: b}

				addressee.ch <- msg
				u.ch <- msg


				break;
			}
		}
	}
}

func (u *User) write(mt int, message []byte) error {
	u.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return u.ws.WriteMessage(mt, message)
}
