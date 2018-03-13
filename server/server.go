package server

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
	"encoding/json"
)

const (
	operatorHandlerPattern = "/api/v1/operator"
	clientHandlerPattern   = "/api/v1/client"
)

type Server struct {
	messages []*Message
	queue    []*Talker
	pairs    map[int]*Pair
	//операции
	//клиент
	addCh   chan *Talker
	delCh   chan *Talker
	queueCh chan int
	//комнаты
	addRoomCh chan map[Talker]Talker
	delRoomCh chan *Talker
	//остальное
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
}

// Create new chat server.
func NewServer() *Server {
	messages := []*Message{}
	queue := []*Talker{}
	pairs := make(map[int]*Pair)
	addCh := make(chan *Talker)
	delCh := make(chan *Talker)
	queueCh := make(chan int)
	addRoomCh := make(chan map[Talker]Talker)
	delRoomCh := make(chan *Talker)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		messages,
		queue,
		pairs,
		addCh,
		delCh,
		queueCh,
		addRoomCh,
		delRoomCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Add(c *Talker) {
	fmt.Println(s.addCh)
	fmt.Println(c)
	s.addCh <- c
	fmt.Println("addCh2")
}

func (s *Server) Del(c *Talker) {
	s.delCh <- c
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) createResponseAllRooms() ResponseMessage {
	response := OperatorResponseRooms{s.pairs, len(s.pairs)}
	jsonstring, _ := json.Marshal(response)
	msg := ResponseMessage{Action: actionGetAllRooms, Status: "OK", Code: 200, Body: jsonstring}
	return msg
}

func (s *Server) checkQueue(t int) {

	fmt.Println(t)
	s.queueCh <- t
	fmt.Println("addCh2")
}

func (s *Server) dispatchPair(pair *Pair) {
	msg := ResponseMessage{Action: initMessage, Status: pair.Talker2.PeerId, Code: 200}
	pair.Talker1.ch <- msg
	msg = ResponseMessage{Action: initMessage, Status: pair.Talker1.PeerId, Code: 200}
	pair.Talker2.ch <- msg
}


// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler for client
	onConnected := func(ws *websocket.Conn) {

		fmt.Println("Talker came")
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		talker := NewTalker(ws, s, nil)
		s.Add(talker)
		s.queue = append(s.queue, talker)
		fmt.Println(len(s.queue))
		s.checkQueue(talker.Id)
		talker.Listen()

	}

	http.Handle(clientHandlerPattern, websocket.Handler(onConnected))
	log.Println("Created handlers")

	for {

		select {

		// Add new a client
		case <-s.addCh:

			// del a client
		case <-s.delCh:
			log.Println("Delete client")

		case <-s.queueCh:
			fmt.Println("Check queue")
			log.Println("Check queue")
			if (len(s.queue) > 1) {
				fmt.Println("get pair")
				pair := NewPair(s)


				pair.Talker1 = s.queue[0]
				pair.Talker1.pair = pair
				s.queue = s.queue[1:]

				pair.Talker2 = s.queue[0]
				pair.Talker2.pair = pair
				s.queue = s.queue[1:]

				s.pairs[pair.Id] = pair
				s.dispatchPair(pair)
			}

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
		fmt.Println("S.addch end")
	}

}
