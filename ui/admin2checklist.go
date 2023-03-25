package main

import (
    "html/template"
    "net/http"
)

type Task struct {
    Title       string
    Description string
}

var tasks []Task

func main() {
    // Serve the admin page
    http.HandleFunc("/", adminHandler)
    // Handle form submission
    http.HandleFunc("/submit", submitHandler)

    // Start the server
    http.ListenAndServe(":8080", nil)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("admin.html")
    t.Execute(w, tasks)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    description := r.FormValue("description")
    task := Task{Title: title, Description: description}
    tasks = append(tasks, task)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
