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
	task, _ := database.GetTask()
	err := index_page.Execute(res, task)
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

	// Prepare an array of finished tasks' ID
	_, task_id := database.GetTask()
	var finished_task_id []int
	for i := 0; i < len(task_id); i++ {
		if req.FormValue("task-"+strconv.Itoa(task_id[i])) == "1" {
			finished_task_id = append(finished_task_id, task_id[i])
		}
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
		database.Write(
			req.FormValue("date"),
			finished_task_id,
			req.FormValue("nippo"),
			task_title,
			new_task,
			deadline,
		)
	}
}

func viewerHandler(res http.ResponseWriter, req *http.Request) {
	target_date := req.FormValue("target_date")

	// Get data in DB
	nippo := database.GetNippo(target_date)

	// Return page information
	viewer_page := template.Must(template.ParseFiles("./root/viewer.html"))
	err := viewer_page.Execute(res, nippo)
	if err != nil {
		panic(err.Error())
	}
}

func serve() {
	// http.Handle("/", http.FileServer(http.Dir("./root/")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))) // Imprt files in assets
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit.html", submitHandler)
	http.HandleFunc("/viewer.html", viewerHandler)
	listen := make(chan bool)
	go func() {
		<-listen
		open.Run("http://localhost:8080/")
		log.Println("browser start")
	}()
	listen <- true
	log.Fatal(http.ListenAndServe(":8080", nil))
}
