package main

import "fmt"
import "net/http"
import "strings"


func main()  {
  fmt.Println("What is your notification's text?")
  var text = ""
  fmt.Scanln(&text)
  fmt.Println("What is your topic's name")
  var topic = ""
  fmt.Scanln(&topic)
  fmt.Println("What is the server's ip/hostname?")
  var host = ""
  fmt.Scanln(&host)

  send(text, topic, host)
}

func send(text string, topic string, host string)  {
  http.Post("http://" + host + "/" + topic, "text/plain",
    strings.NewReader(text))
}
