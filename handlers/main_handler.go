package handlers

import (
    "fmt"
    "net/http"
)

func InitialHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to our simple web server. This is path:%v", r.URL.Path)

    // Initialize the storage service
    storageService := storage.NewStorageService()

    // Get the current state of roommates
    // Create timer 

    

}

func AddRoommateHandler(w http.ResponseWriter, r *http.Request) {
    // Add a new roommate

    storage.AddNewRoommate

}

func AddTaskHandler(w http.ResponseWriter, r *http.Request){
    
}
