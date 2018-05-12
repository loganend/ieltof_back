package interfaces

import (
	"log"
	"github.com/gorilla/websocket"
	"time"
	"net/http"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type Hub struct {
	users      map[uint32]*User
	register   chan *User
	unregister chan *User
	doneCh     chan bool
	errCh      chan error
}

func NewHub() *Hub {

	users := make(map[uint32]*User)
	register := make(chan *User)
	unregister := make(chan *User)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Hub{
		users,
		register,
		unregister,
		doneCh,
		errCh,
	}
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *Hub) Unregister(u *User) {
	h.unregister <- u
}

func (h *Hub) GetUsers() map[uint32]*User {
	return h.users
}

func (s *Hub) Err(err error) {
	s.errCh <- err
}

func (h *Hub) ServeWs(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}



	//user := NewUser(h, ws, u.Id)
	user := NewUser(h, ws, 0)

	//h.register <- user
	user.Listen()
}

func (h *Hub) Listen() {
	log.Println("Listening Hub...")

	//http.HandleFunc("/api/v1/user", h.ServeWs)
	log.Println("Created handlers")

	for {
		select {

		case u := <-h.register:
			h.users[u.Id] = u
			break

		case u := <-h.unregister:

			_, ok := h.users[u.Id]
			if ok {
				delete(h.users, u.Id);
				close(u.ch)
			}
			break

		case err := <-h.errCh:
			log.Println("Error:", err.Error())
			break

		case err := <-h.errCh:
			log.Println("Error:", err.Error())
			break
		case <-h.doneCh:
			break
		}
	}
}
