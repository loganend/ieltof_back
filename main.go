package main


import (
	"log"
	"net/http"
	"github.com/ieltof/server"
)


func main() {
	log.SetFlags(log.Lshortfile)

	server := server.NewServer()
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
