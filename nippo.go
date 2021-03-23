package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/skratchdot/open-golang/open"
	"package.local/database"
)

func main() {
	serve()
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	// Response to the client
	index_page := template.Must(template.ParseFiles("./root/index.html"))
	err := index_page.Execute(res, database.GetTask())
	if err != nil {
		panic(err.Error())
	}
}

func submitHandler(res http.ResponseWriter, req *http.Request) {
	// Response to the client
	submit_page := template.Must(template.ParseFiles("./root/submit.html"))
	err := submit_page.Execute(res, nil)
	if err != nil {
		panic(err.Error())
	}

	// Prepare an array of new tasks
	var number_of_new_task int
	number_of_new_task, _ = strconv.Atoi(req.FormValue("number-of-new-tasks"))
	var task_title []string
	var new_task []string
	var deadline []string
	for i := 1; i <= number_of_new_task; i++ {
		task_title = append(task_title, req.FormValue("title-"+strconv.Itoa(i)))
		new_task = append(new_task, req.FormValue("new-task-"+strconv.Itoa(i)))
		deadline = append(deadline, req.FormValue("deadline-"+strconv.Itoa(i)))
	}

	// DB operation
	if req.FormValue("date") != "" {
		var arr []string = []string{"Sample", "Values"}
		database.Write(
			req.FormValue("date"),
			arr,
			req.FormValue("nippo"),
			task_title,
			new_task,
			deadline,
		)
	}
}

func serve() {
	// http.Handle("/", http.FileServer(http.Dir("./root/")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) // Imprt files in assets
	http.HandleFunc("/", indexHandler)
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
