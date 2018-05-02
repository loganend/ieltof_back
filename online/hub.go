package online

import (
	"log"
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"github.com/ieltof/interfaces"
	"github.com/ieltof/domain"
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
	handler    interfaces.WebserviceHandler
}

func NewHub(handler interfaces.WebserviceHandler) *Hub {

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
		handler,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
}

func (h *Hub) Unregister(u *User) {
	h.unregister <- u
}

func (s *Hub) Err(err error) {
	s.errCh <- err
}

func (h *Hub) serveWs(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Println(err)
		return
	}

	cookie, err := req.Cookie("facebookid")
	var facebookId = cookie.Value
	if err != nil {
		log.Println(err)
		return
	}

	var udb domain.User
	udb.FacebookId = facebookId
	u, err := h.handler.Interator.GetUser(udb);
	if err != nil {
		log.Println(err)
		return
	}

	user := NewUser(h, ws, u.Id)
	h.register <- user
	user.Listen()
}

func (h *Hub) Listen() {
	log.Println("Listening Hub...")

	http.HandleFunc("/api/v1/user", h.serveWs)
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
