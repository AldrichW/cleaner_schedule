// main.go
// ***STARTING POINT*** of backend. Initialize all components and subservices here
package main

import (
    "fmt"
    "net/http" 
	"aldrich_projects/cleaner_schedule/handlers"
)


func initialHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World!")
    //Call on the handlers package
    handlers.InitialHandler(w, r)

}

func addRoommateHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "[main] Roommate Handler")
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "[main] Task Handler")
}

func main () {
    
    fmt.Println("Welcome to Task Roommate")

    //Initialize datastore service

    // storageService := storage.NewStorageService()
    // if storage_service == nil {
    //     // Oops, something went wrong with initialization
    //     fmt.Errorf("Error initializing storage service")
    // }

    // Start server. Pass in nil for now
    http.HandleFunc("/", initialHandler)
    http.HandleFunc("/add/roommate/", addRoommateHandler)
    http.HandleFunc("/add/task/", addTaskHandler)
    http.ListenAndServe(":9300", nil)
}