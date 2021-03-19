package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/skratchdot/open-golang/open"
	"package.local/database"
)

func main() {
	serve()
}

func submitHandler(res http.ResponseWriter, req *http.Request) {
	nippo := req.FormValue("nippo")
	log.Println(nippo)
	fmt.Fprint(res, readFile("submit.html"))
	// DB operation
	if req.FormValue("date") != "" {
		var arr []string = []string{"Sample", "Values"}
		database.Write(
			req.FormValue("date"),
			arr,
			req.FormValue("nippo"),
			arr,
			arr,
		)
	}
}

// Test function of html/template
func templateHandler(res http.ResponseWriter, req *http.Request) {
	var tplt *template.Template
	tplt = template.Must(template.ParseFiles("./root/template.html"))
	age := 256
	err := tplt.Execute(res, age)
	if err != nil {
		panic(err.Error())
	}
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
	http.HandleFunc("/template", templateHandler)
	listen := make(chan bool)
	go func() {
		<-listen
		open.Run("http://localhost:8080/")
		log.Println("browser start")
	}()
	listen <- true
	log.Fatal(http.ListenAndServe(":8080", nil))
}
