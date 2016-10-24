package main

import (
  "net/http"
  "io"
  "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "hello, world!\n")
}

func main() {
  http.HandleFunc("/Hello", HelloServer)
  log.Fatal(http.ListenAndServe(":420", nil))
}
