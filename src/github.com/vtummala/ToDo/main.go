package main

//Packages to be imported
import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "log"
)

//Main function
func main() {

    /* Create a Multiplexer to which matches the incoming requests
       againt the available, registered routes*/
    router := mux.NewRouter()

    // Routes handles and methods for EndPoints /tasks, /tasks/{id}, /tasks/{id}
    router.HandleFunc("/tasks", GetTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}", AddTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
    fmt.Printf("hello, world\n")
}

func GetTasks(w http.ResponseWriter, r *http.Request){}
func AddTask(w http.ResponseWriter, r *http.Request){}
func DeleteTask(w http.ResponseWriter, r *http.Request){}
