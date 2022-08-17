package main

import (
   "net/http"
)





func main() {
    router := http.NewServeMux()
    router.HandleFunc("/event")
    router.HandleFunc("/time")
}
