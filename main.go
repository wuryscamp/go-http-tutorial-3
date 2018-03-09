package main

import (
  "fmt"
  "net/http"
  "log"
)

 // req 1 => req 2

 func logRequest(h http.Handler) http.Handler {
   return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
     log.Printf("%s request %s", req.RemoteAddr, req.URL)
     h.ServeHTTP(res, req)
   })
 }


func main() {

  h := http.NewServeMux()// http.Handler (interface)

  h.HandleFunc("/me", func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Me Handler")
  })

  h.HandleFunc("/you", func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "You Handler")
  })

  h.HandleFunc("/her", func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Her Handler")
  })

  h.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Index Handler")
  })

  // logger middleware
  h1 := logRequest(h)

  log.Println("Server running on port 3000")

  //
  err := http.ListenAndServe(":3000", h1)

  if err != nil{
    log.Fatal(err)
  }


}
