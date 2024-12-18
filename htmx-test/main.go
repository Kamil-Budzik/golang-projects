package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var tmpl *template.Template
var tasks []Task

type Task struct {
	Id   int
	Task string
	Done bool
}

func init() {
	tmpl, _ = template.ParseGlob("templates/*.html")
}

func main() {
	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", Homepage)
	gRouter.HandleFunc("/tasks", fetchTasks).Methods("GET")
	gRouter.HandleFunc("/newtaskform", getTaskForm)
	gRouter.HandleFunc("/tasks", addTask).Methods("POST")
	gRouter.HandleFunc("/gettaskupdateform/{id}", getTaskUpdateForm).Methods("GET")
	gRouter.HandleFunc("/tasks/{id}", updateTask).Methods("PUT", "POST")
	gRouter.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Println("Server is running on http://localhost:4000")
	http.ListenAndServe(":4000", gRouter)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func fetchTasks(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "todoList", tasks)
}

func getTaskForm(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "addTaskForm", nil)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	newTask := Task{
		Id:   len(tasks) + 1,
		Task: task,
		Done: false,
	}
	fmt.Println(newTask)
	tasks = append(tasks, newTask)
	tmpl.ExecuteTemplate(w, "todoList", tasks)
}

func getTaskUpdateForm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, _ := strconv.Atoi(vars["id"])
	for _, task := range tasks {
		if task.Id == taskId {
			tmpl.ExecuteTemplate(w, "updateTaskForm", task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, _ := strconv.Atoi(vars["id"])
	for i, task := range tasks {
		if task.Id == taskId {
			task.Task = r.FormValue("task")
			task.Done = strings.ToLower(r.FormValue("done")) == "yes" || strings.ToLower(r.FormValue("done")) == "on"
			tasks[i] = task
			break
		}
	}
	tmpl.ExecuteTemplate(w, "todoList", tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, _ := strconv.Atoi(vars["id"])
	for i, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	tmpl.ExecuteTemplate(w, "todoList", tasks)
}
