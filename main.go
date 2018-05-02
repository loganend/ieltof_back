package main


import (
	"log"
	"os"
	"net/http"
	"github.com/ieltof/route"
	"github.com/ieltof/server"
	"github.com/ieltof/infrastructure"
	"github.com/ieltof/shared/jsonconfig"
	"github.com/ieltof/interfaces"
	"github.com/ieltof/usecases"
	"encoding/json"
	"github.com/ieltof/online"
)


func main() {
	log.SetFlags(log.Lshortfile)

	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	dbHandler := infrastructure.NewPostgresHandler(config.Database)

	handlers := make(map[string] interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbFriendRepo"] = dbHandler
	handlers["DbMessageRepo"] = dbHandler

	//userInteractor := new(usecases.UserInteractor)
	//friendInterator := new(usecases.FriendInteractor)
	//messageInterator := new(usecases.MessageInterator)
	interator := new(usecases.Interactor);

	interator.MessageRepository = interfaces.NewDbMessageRepo(handlers)
	interator.FriendRepository = interfaces.NewDbFriendRepo(handlers)
	interator.UserRepository = interfaces.NewDbUserRepo(handlers)


	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.Interator = interator
	//webserviceHandler.UserInteractor = userInteractor
	//webserviceHandler.FriendInteractor = friendInterator
	//webserviceHandler.MessageInteractor = messageInterator

	hub := online.NewHub(webserviceHandler);
	go hub.Listen()

	server := server.NewServer()
	go server.Listen()

	 //static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":8080", route.LoadHTTP(webserviceHandler)))
}

var config = &configuration{}

type configuration struct {
	Database  infrastructure.Info   `json:"Database"`
	Server    server.Server   `json:"Server"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
