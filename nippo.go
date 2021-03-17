package main

import (
	"log"
	"net/http"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	serve()
}

func submitHandler(res http.ResponseWriter, req *http.Request) {
	nippo := req.FormValue("nippo")
	log.Println(nippo)
	// DB operation
}

// Ref: https://gist.github.com/niratama/6b0117c6c6f2d21b5687
func serve() {
	fs := http.FileServer(http.Dir("./static/html/"))
	http.Handle("/", fs)
	http.HandleFunc("/submit.html", submitHandler)
	listen := make(chan bool)
	go func() {
		<-listen
		open.Run("http://localhost:8080/")
		log.Println("browser start")
	}()
	listen <- true
	log.Fatal(http.ListenAndServe(":8080", nil))
}
