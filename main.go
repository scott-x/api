package main

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/scott-x/api/handler"
)

func RegisterHandlers() *httprouter.Router {
  router := httprouter.New()
  router.POST("/user", handler.CreateUser)
  router.POST("/user/:username", handler.Login)
  return router
}

func main() {
  r := RegisterHandlers()
  fmt.Println("server is running at https://127.0.0.1:8000")
  http.ListenAndServe(":8000", r)
}
