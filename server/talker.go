package server

import (
	"golang.org/x/net/websocket"
	"log"
	"io"
	"fmt"
	"encoding/json"
	"time"
)

const channelBufSize = 100

var maxId int = 0

type Talker struct {
	Id        int    `json:"id"`
	PeerId    string `json:"peerId"`
	ready     bool
	testId    int
	ws        *websocket.Conn
	server    *Server
	pair      *Pair
	ch        chan ResponseMessage
	doneCh    chan bool
	addRoomCh chan *Pair
	delRoomCh chan *Pair
}

func NewTalker(ws *websocket.Conn, server *Server, pair *Pair) *Talker {

	if ws == nil {
		panic("ws cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	maxId++
	ch := make(chan ResponseMessage, channelBufSize)
	doneCh := make(chan bool)
	addRoomCh := make(chan *Pair)
	delRoomCh := make(chan *Pair)
	return &Talker{maxId, "", false, 9999, ws, server, pair, ch, doneCh, addRoomCh, delRoomCh}
}

func (c *Talker) Conn() *websocket.Conn {
	return c.ws
}

func (c *Talker) Done() {
	c.doneCh <- true
}

func (c *Talker) Listen() {

	go c.listenWrite()
	c.listenRead()
}

func (c *Talker) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			//log.Println("Send:", msg, c.ws)
			websocket.JSON.Send(c.ws, msg)

			// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return

			// case r := <-c.addRoomCh:
			// 	log.Println("add room to client")
			// 	c.room = r
			// 	msg := ResponseMessage{Action: actionCreateRoom, Status: "OK", Code: 200}
			// 	websocket.JSON.Send(c.ws, msg)
			//
			// msg := ClientGreetingResponse{"grabing", "ROOM create hello:)"}
			// websocket.JSON.Send(c.ws, msg)

		}

	}
}

func (t *Talker) dispatchDisconectMessage(t_pair *Talker) {

	t_pair.pair = nil;
	msg := ResponseMessage{Action: disconectMessage, Status: t_pair.PeerId, Code: 200}
	t_pair.ch <- msg

}

// Listen read request via chanel
func (c *Talker) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenWrite method
			return

			// read data from websocket connection
		default:
			var msg RequestMessage
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			}
			//log.Println(msg)

			switch msg.Action {

			case actionGetTest:

				var test Test;
				if(c.testId == 9999) {
					c.testId = c.Id;
				}else{
					c.testId = c.testId + 1;
				}

				switch c.testId % 6 {
				case 1:
					test = getTest1()
				case 2:
					test = getTest2()
				case 3:
					test = getTest3()
				case 4:
					test = getTest4()
				case 5:
					test = getTest5()
				case 0:
					test = getTest6()
				}

				b, err := json.Marshal(test)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(test)
				fmt.Println(b)
				c.ch <- ResponseMessage{Action: actionGetTest, Status: "OK", Code: 200, Body: b, Test: test}

			case actionToken:
				log.Println(actionToken)
				var token TokenBody
				err := json.Unmarshal(msg.Body, &token)
				if !CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := ResponseMessage{Action: actionToken, Status: "Invalid Request", Code: 403}
					c.ch <- msg
				}
				log.Println(token.Token)

				c.PeerId = token.Token;

			case initViedoChat:

				log.Println(initViedoChat)
				var message Message
				err := json.Unmarshal(msg.Body, &message)
				if !CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := ResponseMessage{Action: initViedoChat, Status: "Invalid Request", Code: 403}
					c.ch <- msg
				}
				c.ready = true;
				if (len(c.server.queue) > 1) {
					pair := NewPair(c.server)
					log.Println("queue len")
					log.Println(len(c.server.queue))
					for i := 0; i < len(c.server.queue); i++ {
						log.Println(i);
						log.Println("c.server.queue[i].Id %d \n", c.server.queue[i].Id)
						log.Println("c.Id %d \n", c.Id)
						log.Println("c.server.queue[i].ready %d \n", c.server.queue[i].ready)
						if (c.server.queue[i].Id != c.Id && c.server.queue[i].ready) {
							pair.Talker1 = c.server.queue[i]
							pair.Talker1.pair = pair
							pair.Talker1.ready = false;
							//c.server.queue = c.server.queue[i+1:]

							pair.Talker2 = c;
							pair.Talker2.pair = pair
							c.ready = false;

							//c.server.queue = c.server.queue[i+1:]

							c.server.pairs[pair.Id] = pair
							c.server.dispatchPair(pair)
						}
					}
				}

			case getNextPartner:

				log.Println(getNextPartner)
				var message Message
				err := json.Unmarshal(msg.Body, &message)
				if !CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := ResponseMessage{Action: getNextPartner, Status: "Invalid Request", Code: 403}
					c.ch <- msg
				}

				if (c.pair != nil) {

					log.Println(c.Id);
					log.Println(c.pair.Talker1.Id);
					log.Println(c.pair.Talker2.Id);
					if (c.pair.Talker1.Id != c.Id) {
						c.dispatchDisconectMessage(c.pair.Talker1);
					} else {
						c.dispatchDisconectMessage(c.pair.Talker2);
					}
					c.pair = nil;
				}

				c.ready = true;
				if (len(c.server.queue) > 1) {
					pair := NewPair(c.server)
					log.Println("queue len")
					log.Println(len(c.server.queue))
					for i := 0; i < len(c.server.queue); i++ {
						log.Println(i);
						log.Println("c.server.queue[i].Id %d \n", c.server.queue[i].Id)
						log.Println("c.Id %d \n", c.Id)
						log.Println("c.server.queue[i].ready %d \n", c.server.queue[i].ready)
						if (c.server.queue[i].Id != c.Id && c.server.queue[i].ready) {
							pair.Talker1 = c.server.queue[i]
							pair.Talker1.ready = false;
							pair.Talker1.pair = pair

							//c.server.queue = c.server.queue[i+1:]

							pair.Talker2 = c;
							pair.Talker2.pair = pair
							c.ready = false;

							//c.server.queue = c.server.queue[i+1:]

							c.server.pairs[pair.Id] = pair
							c.server.dispatchPair(pair)
						}
					}
				}

			case stopVideoChat:

				log.Println(stopVideoChat)
				var message Message
				err := json.Unmarshal(msg.Body, &message)
				if !CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := ResponseMessage{Action: stopVideoChat, Status: "Invalid Request", Code: 403}
					c.ch <- msg
				}

				if (c.pair != nil) {
					log.Println(c.Id);
					log.Println(c.pair.Talker1.Id);
					log.Println(c.pair.Talker2.Id);
					if (c.pair.Talker1.Id != c.Id) {
						c.dispatchDisconectMessage(c.pair.Talker1);
					} else {
						c.dispatchDisconectMessage(c.pair.Talker2);
					}
					c.pair = nil;
				}
				c.ready = false;

				//отправка сообщений
			case actionSendMessage:

				log.Println(actionSendMessage)
				var message Message
				err := json.Unmarshal(msg.Body, &message)
				if !CheckError(err, "Invalid RawData"+string(msg.Body), false) {
					msg := ResponseMessage{Action: actionSendMessage, Status: "Invalid Request", Code: 403}
					c.ch <- msg
				}

				if(c.pair != nil) {
					if (c.pair.Talker2 != nil) {
						if (c.pair.Talker1 == c) {
							message.Author = "client"
							message.Room = c.pair.Id
							message.Time = int(time.Now().Unix())
							b, err := json.Marshal(message)
							if err != nil {
								fmt.Println(err)
								return
							}
							c.pair.Talker2.ch <- ResponseMessage{Action: actionSendMessage, Status: message.Body, Code: 200, Body: b}
							//c.pair.channelForMessage <- message
						} else {
							message.Author = "client"
							message.Room = c.pair.Id
							message.Time = int(time.Now().Unix())
							b, err := json.Marshal(message)
							if err != nil {
								fmt.Println(err)
								return
							}
							c.pair.Talker1.ch <- ResponseMessage{Action: actionSendMessage, Status: message.Body, Code: 200, Body: b}
						}
					} else {
						msg := ResponseMessage{Action: actionSendMessage, Status: "Room not found", Code: 404}
						c.ch <- msg
					}
				}

			case actionCloseRoom:
				log.Println(actionCloseRoom)
				c.pair.Status = roomClose
				c.pair.channelForStatus <- roomClose

			case actionGetAllMessages:
				log.Println(actionGetAllMessages)
				messages, _ := json.Marshal(c.pair.Messages)
				response := ResponseMessage{Action: actionGetAllMessages, Status: "OK", Code: 200, Body: messages}
				log.Println(response)
				c.ch <- response
			}
		}
	}
}

//case initMessage:
//
//
//log.Println(initMessage)
//var message Message
//err := json.Unmarshal(msg.Body, &message)
//if !CheckError(err, "Invalid RawData"+string(msg.Body), false) {
//msg := ResponseMessage{Action: initMessage, Status: "Invalid Request", Code: 403}
//c.ch <- msg
//}
//
//if (c.pair.Talker2 != nil) {
//if (c.pair.Talker1 == c) {
//message.Author = c.PeerId
//message.Room = c.pair.Id
//message.Time = int(time.Now().Unix())
//b, err := json.Marshal(message)
//if err != nil {
//fmt.Println(err)
//return
//}
//c.pair.Talker2.ch <- ResponseMessage{Action: initMessage, Status: "OK", Code: 200, Body: b}
////c.pair.channelForMessage <- message
//} else {
//message.Author =  c.PeerId
//message.Room = c.pair.Id
//message.Time = int(time.Now().Unix())
//b, err := json.Marshal(message)
//if err != nil {
//fmt.Println(err)
//return
//}
//c.pair.Talker1.ch <- ResponseMessage{Action: initMessage, Status: "OK", Code: 200, Body: b}
//}
//} else {
//msg := ResponseMessage{Action: initMessage, Status: "Room not found", Code: 404}
//c.ch <- msg
//}
