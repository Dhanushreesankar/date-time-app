package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Load the local time zone (change "Asia/Kolkata" if needed)
    loc, _ := time.LoadLocation("Asia/Kolkata")
    currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05 MST")
    fmt.Fprintf(w, "Current Date & Time: %s", currentTime)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at :8080")
    http.ListenAndServe(":8080", nil)
}
