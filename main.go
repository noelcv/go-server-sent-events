package main

import (
   "net/http"
   "fmt"
)

var msgChannel chan string

func sseHandler(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Access-Control-Allow-Origin", "*") //cors
   w.Header().Set("Content-Type", "text/event-stream")
   w.Header().Set("Cache-Control", "no-cache")
   w.Header().Set("Connection", "keep-alive")
   
   
   msgChannel = make(chan string)
   
   //handle client closing channel
   defer func() {
      close(msgChannel)
      msgChannel = nil
      fmt.Println("Client closed connection")
   }()
   
   flusher, ok := w.(http.Flusher)
   if !ok {
      fmt.Println("Couldn't init http.Flusher")
   }
   
   for {
      select {
        case message := <- msgChannel:
         fmt.Fprintf(w, "data: %s \n\n", message)
         flusher.Flush()
        case <- r.Context().Done():
         fmt.Println("Client closed connection")
         return
      }
   }
   
}


func main() {
    router := http.NewServeMux()
    router.HandleFunc("/event", sseHandler)
   //  router.HandleFunc("/time", )
}
