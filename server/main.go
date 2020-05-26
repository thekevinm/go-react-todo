package main

import (
    "fmt"
    "log"
    "net/http"
    "./router"
)

func main() {
    // create an instance of `router`
    r := router.Router()
    // serve on port 8080
    fmt.Println("Starting server on the port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
