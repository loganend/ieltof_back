package route

import (
	"net/http"

	hr "github.com/ieltof/route/httprouterwrapper"
	"github.com/ieltof/route/logrequest"
	"github.com/ieltof/interfaces"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/ieltof/server"
	"golang.org/x/net/websocket"
)

func LoadHTTP(webserviceHandler interfaces.WebserviceHandler, hub *interfaces.Hub, server *server.Server) http.Handler {
	return middleware(routes(webserviceHandler, hub, server))
}

func routes(webserviceHandler interfaces.WebserviceHandler, hub *interfaces.Hub, server *server.Server) *httprouter.Router {
	r := httprouter.New()
	//r.NotFound = alice.
	//	New().ThenFunc(controller.Error404)

	r.POST("/api2/user",  hr.Handler(alice.New().ThenFunc(webserviceHandler.GetUser)))
	r.GET("/api2/users", hr.Handler(alice.New().ThenFunc(webserviceHandler.GetUsers)))
	r.GET("/api2/users/online", hr.Handler(alice.New().ThenFunc(webserviceHandler.GetOnlineUsers)))

	r.POST("/api2/friend",  hr.Handler(alice.New().ThenFunc(webserviceHandler.FriendRequest)))
	r.GET("/api2/friends", hr.Handler(alice.New().ThenFunc(webserviceHandler.GetFriends)))

	r.PUT("/api/friend/accept", hr.Handler(alice.New().ThenFunc(webserviceHandler.AcceptFriendship)))
	r.DELETE("/api/friend/ignore", hr.Handler(alice.New().ThenFunc(webserviceHandler.IgnoreFriendship)))


	r.OPTIONS("/api/friend/accept", hr.Handler(alice.New().ThenFunc(webserviceHandler.OptionRequest)))
	r.OPTIONS("/api/friend/ignore", hr.Handler(alice.New().ThenFunc(webserviceHandler.OptionRequest)))
	r.OPTIONS("/api2/friend", hr.Handler(alice.New().ThenFunc(webserviceHandler.OptionRequest)))

	r.GET("/api/v1/user", hr.Handler(alice.New().ThenFunc(hub.ServeWs)))
	r.GET("/api/v1/client", hr.Handler(websocket.Handler(server.InitVideoChat)));

	return r
}

func middleware(h http.Handler) http.Handler {

	h = logrequest.Handler(h)
	h = context.ClearHandler(h)

	return h
}
