package main

import (
  "fmt"
  "net/http"
)

func GetType(url string)(string,error){
  content, err := http.Get(url)
  if err != nil{
    return "",err
  }
  defer content.Body.Close()

  content_type := content.Header.Get("content-type")

  if content_type == ""{
    return "",fmt.Errorf("Cannot find content-type")
  }

  return content_type, nil
}

func main() {
  ctype, err := GetType("https://linkedin.com")

  if err != nil {
    fmt.Printf("ERROR: %v\n", err)
  } else {
    fmt.Printf("content-type: %v\n", ctype)
  }
}