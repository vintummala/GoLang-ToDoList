package main

//Packages to be imported
import (
    "fmt"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "math/rand"
    "strconv"
    "log"
)

//Task Structs (Model)

type Task struct {
  ID		string    `json:"id"`
  Title		string    `json:"title"`
//  CreatedAt	time.Time `json:"created_at"`
  Completed	bool      `json:"completed"`
}

//Init Tasks Variable as a slice Task Struct
var tasks []Task

//Main function
func main() {

    /* Create a Multiplexer which matches the incoming requests
       againt the available, registered routes*/
    router := mux.NewRouter()

    //Statis Data
    tasks = append(tasks, Task{ID: "1", Title: "Go Challange", Completed: true})
    tasks = append(tasks, Task{ID: "2", Title: "Go REST Api", Completed: true})
    tasks = append(tasks, Task{ID: "3", Title: "Go Mock database", Completed: true})

    // Route handlers and methods for EndPoints /tasks, /tasks/{id}, /tasks/{id}
    router.HandleFunc("/tasks", GetTasks).Methods("GET")
    router.HandleFunc("/tasks", AddTask).Methods("POST")
    router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

    //Run the Server
    log.Fatal(http.ListenAndServe(":8000", router))
    fmt.Printf("hello, world\n")
}

//Get All Tasks
func GetTasks(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-type", "application/json")
   json.NewEncoder(w).Encode(tasks)
}

//Add/Create a New Task
func AddTask(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-type", "application/json")
   var task Task

   //Decodes the data in body of the request and copies it to the task struct
   _ = json.NewDecoder(r.Body).Decode(&task)

   //Creating a random task ID
   task.ID = strconv.Itoa(rand.Intn("10000000"))

   //append to the main tasks slice
   tasks = append(tasks, task)

   //Output the task or write to an output stream or kind of printf
   json.NewEncoder(w).Encode(task)
}

//Delete a task
func DeleteTask(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-type", "application/json")

    //Get the ID using mux.Vars
    params := mux.Vars(r)

   for index, item := range tasks {
       if item.ID == params["id"] {
          //delets the task with that index
          tasks = append(tasks[:index],tasks[index+1:]...)
          break
      }
   }

   json.NewEncoder(w).Encode(books)
}
