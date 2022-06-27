package main

import (
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func rootHandler(resp http.ResponseWriter, req *http.Request) {
  resp.Write([]byte(`{"msg":"Succesfully reached site!"}`))
}

func main() {
  mainRouter := mux.NewRouter()

  mainRouter.HandleFunc("/", rootHandler).Methods(http.MethodGet)

  log.Fatal(http.ListenAndServe("localhost:8000", mainRouter))
}
