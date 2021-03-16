package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	serve()
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, world.")
}

// Ref: https://gist.github.com/niratama/6b0117c6c6f2d21b5687
func serve() {
	http.HandleFunc("/", handler)
	listen := make(chan bool)
	go func() {
		<-listen
		open.Run("http://localhost:8080/")
		log.Println("browser start")
	}()
	listen <- true
	log.Fatal(http.ListenAndServe(":8080", nil))
}
