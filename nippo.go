package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Fprint(res, readFile("submit.html"))
	// DB operation
}

// Read any files in ./root/
func readFile(fileName string) string {
	bytes, err := ioutil.ReadFile("./root/" + fileName)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func serve() {
	http.Handle("/", http.FileServer(http.Dir("./root/")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) // Imprt files in assets
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
