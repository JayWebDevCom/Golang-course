package main

import (
"net/http"
"io/ioutil"
"fmt"
)

func httpStuff() {
  // blank identifier avoids issue of an unused variable
  // in this case an error
  res, _ := http.Get("http://www.example.com")
  page, _ := ioutil.ReadAll(res.Body)
  res.Body.Close()

  // print as a string
  fmt.Printf("%s", page)
}
