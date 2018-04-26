package route

import (
	"net/http"

	hr "github.com/ieltof/route/httprouterwrapper"
	"github.com/ieltof/route/logrequest"
	"github.com/ieltof/interfaces"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func LoadHTTP(webserviceHandler interfaces.WebserviceHandler) http.Handler {
	return middleware(routes(webserviceHandler))
}

func routes(webserviceHandler interfaces.WebserviceHandler) *httprouter.Router {
	r := httprouter.New()

	//r.NotFound = alice.
	//	New().ThenFunc(controller.Error404)

	r.POST("/api/user",  hr.Handler(alice.New().ThenFunc(webserviceHandler.GetUser)))
	r.GET("/api/users", hr.Handler(alice.New().ThenFunc(webserviceHandler.GetUsers)))

	r.POST("/api/friend",  hr.Handler(alice.New().ThenFunc(webserviceHandler.FriendRequest)))
	r.GET("/api/friends", hr.Handler(alice.New().ThenFunc(webserviceHandler.GetFriends)))

	return r
}

func middleware(h http.Handler) http.Handler {

	h = logrequest.Handler(h)
	h = context.ClearHandler(h)

	return h
}
