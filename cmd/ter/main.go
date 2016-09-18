package main

import (
  "os"
  "runtime"
  "fmt"
  _ "log"
  "net/http"
  "io/ioutil"
  app "github.com/callowaylc/ter"

)

func init() {
  if cpu := runtime.NumCPU(); cpu == 1 {
    runtime.GOMAXPROCS(2)
  } else {
    runtime.GOMAXPROCS(cpu)
  }
}


func main() {
  config := app.GetConfig()
  message := fmt.Sprintf("start application with config %+v", config)
  app.Logs(message, nil)


  // get list and parse
  content, err := app.Memoize(func() interface{} {
    resp, err := http.Get(config.List)
    if err != nil {
      // handle error
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    return string(body)
  });

  if err != nil {

  }

  fmt.Println(content)

  /*
  response, err := http.Get("http://www.theeroticreview.com")
  if err != nil {
    log.Fatal(err)
  }
  defer response.Body.Close()

  fmt.Println("response Status:", response.Status)
  fmt.Println("response Headers:", response.Header)
  body, _ := ioutil.ReadAll(response.Body)
  fmt.Println("response Body:", string(body))
  */
  os.Exit(0)
}
