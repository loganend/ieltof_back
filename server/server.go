package server

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
)

const (
	operatorHandlerPattern = "/api/v1/operator"
	clientHandlerPattern   = "/api/v1/client"
)

type Server struct {
	messages []*Message
	pairs     map[int]*Pair
	//операции
	//клиент
	addCh chan *Talker
	delCh chan *Talker

	//комнаты
	//addRoomCh chan map[Client]Operator
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
	//operators := make(map[int]*Operator)
	pairs := make(map[int]*Pair)
	addCh := make(chan *Talker)
	delCh := make(chan *Talker)
	//addOCh := make(chan *Operator)
	//delOCh := make(chan *Operator)
	addRoomCh := make(chan map[Talker]Talker)
	delRoomCh := make(chan *Talker)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		messages,
		//operators,
		pairs,
		addCh,
		delCh,
		//addOCh,
		//delOCh,
		addRoomCh,
		delRoomCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}


func (s *Server) Add(c *Talker) {
	s.addCh <- c
}

//func (s *Server) AddOperator(o *Operator) {
//	s.addOCh <- o
//}

func (s *Server) Del(c *Talker) {
	s.delCh <- c
}

//func (s *Server) DelOperator(o *Operator) {
//	s.delOCh <- o
//}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

//func (s *Server) broadcast(responseMessage ResponseMessage) {
//	for _, operator := range s.operators {
//		operator.ch <- responseMessage
//	}
//}

//func (s *Server) createResponseAllRooms() ResponseMessage {
//	response := OperatorResponseRooms{s.rooms, len(s.rooms)}
//	jsonstring, _ := json.Marshal(response)
//	msg := ResponseMessage{Action: actionGetAllRooms, Status: "OK", Code: 200, Body: jsonstring}
//	return msg
//}

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

		pair := NewPair(s)
		talker := NewTalker(ws, s, pair)
		pair.Talker = talker
		s.pairs[pair.Id] = pair
		s.Add(talker)
		pair.Listen()
		talker.Listen()
	}

	// websocket handler for operator
	onConnectedOperator := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		//operator := NewOperator(ws, s)
		//s.AddOperator(operator)
		//operator.Listen()
	}
	http.Handle(clientHandlerPattern, websocket.Handler(onConnected))
	http.Handle(operatorHandlerPattern, websocket.Handler(onConnectedOperator))
	log.Println("Created handlers")

	for {
		select {

		//// Add new a client
		//case <-s.addCh:
		//	msg := s.createResponseAllRooms()
		//	s.broadcast(msg)
		//
		//	// del a client
		//case <-s.delCh:
		//	log.Println("Delete client")
		//	msg := s.createResponseAllRooms()
		//	s.broadcast(msg)

		//	// Add new a operator
		//case o := <-s.addOCh:
		//	log.Println("Added new operator")
		//	s.operators[o.Id] = o
		//	msg := s.createResponseAllRooms()
		//	o.ch <- msg
		//
		//	// del a operator
		//case o := <-s.delOCh:
		//	log.Println("Delete operator")
		//	delete(s.operators, o.Id)
		//
		//case err := <-s.errCh:
		//	log.Println("Error:", err.Error())
		//
		//case <-s.doneCh:
		//	return
		//}
		}
	}

}
