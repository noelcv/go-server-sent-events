package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var msgChannel chan string

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Access-Control-Allow-Origin", "*") //cors
	if msgChannel != nil {
      msg := time.Now().Format("15:04:05")
		msgChannel <- msg
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Println("Client connected")
	w.Header().Set("Access-Control-Allow-Origin", "*") //cors
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
   
   //create a channel
	msgChannel = make(chan string)

	//handle client closing channel
	defer func() {
		close(msgChannel)
		msgChannel = nil
		fmt.Println("Client closed connection")
	}()
   
   //clean the Buffered data
	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("Couldn't init http.Flusher")
	}

	for {
		select {
      /*like switch statements but for channels: write message in a stringified object
      and flush afterwards*/
		case message := <- msgChannel:
			fmt.Fprintf(w, "data: %s\n\n", message)
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
	router.HandleFunc("/time", getTimeHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
