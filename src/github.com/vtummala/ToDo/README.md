Go Installation: https://golang.org/doc/install?download=go1.10.2.linux-amd64.tar.gz

1. The executible generated will be the name of the folder in which .go files are present.
2. Steps for compiling and building go files is
	go build
	./<exe name>
3. Start the file with package main
4. If there are multiple packages to be imported, the syntax is as below
import (
    package 1
	package 2
)
5. The name mux stands for "HTTP request multiplexer". Like the standard http.ServeMux, mux.
Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions.

mux stands for HTTP request multiplexer which matches an incoming request to against a list of routes (registered)

The package to be imported is "github.com/gorilla/mux"
3. log.Fatal(http.ListenAndServe(":8000", router))
Package to be imported for this api to be used in go lang is "net/http"
Since we are having log.Fatal, the package to be imported for logging is "log"
ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux.
Handle and HandleFunc add handlers to DefaultServeMux:


4. Github pushing commands followed:
https://help.github.com/articles/adding-an-existing-project-to-github-using-the-command-line/

5.ToDoList google spreadsheet:  https://docs.google.com/spreadsheets/d/1ycwCp3b9d15d_r9D4G7FFXeezbowVHHXQjhjPYkKfCg/edit#gid=0

References:

https://www.youtube.com/watch?v=2KmHtgtEZ1s
https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
http://www.gorillatoolkit.org/pkg/mux
https://medium.com/@thedevsaddam/lets-make-a-simple-todo-app-with-go-1d5998acbc15
https://developers.google.com/sheets/api/quickstart/go
https://developers.google.com/fusiontables/docs/v2/getting_started#invoking -- important
https://developers.google.com/sheets/api/quickstart/go

https://medium.com/@maumribeiro/a-fullstack-epic-part-i-a-rest-api-in-go-accessing-mongo-db-608b46e969cd : This is using mongoDB. Try this 
https://medium.com/@thedevsaddam/lets-make-a-simple-todo-app-with-go-1d5998acbc15 : This uses Chi-router. Instead we can use gorilla-Mux as well.

https://www.youtube.com/watch?v=SonwZ6MF5BE
https://github.com/vintummala/GoLang-ToDoList.git
