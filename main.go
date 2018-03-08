package main


import (
	"log"
	"net/http"
	"github.com/ieltof/server"
	"fmt"
)


func main() {
	fmt.Println("hell")
	log.SetFlags(log.Lshortfile)

	fmt.Println("hell")
	server := server.NewServer()
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
