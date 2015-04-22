package main

import (
  "fmt"
  "net/http"
  "os"
  "io/ioutil"
  "github.com/rs/cors"
)

func main() {

  h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    data, err := ioutil.ReadFile("./data.json")
    fmt.Fprintln(w, string(data))
    if err != nil {
      panic(err)
    }
  })

  // cors.Default() setup the middleware with default options being
  // all origins accepted with simple methods (GET, POST). See
  // documentation below for more options.
  handler := cors.Default().Handler(h)
  http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
