package main

import (
  "os"
  "runtime"
  "fmt"
  _ "log"
  app "github.com/callowaylc/zd"

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


  app.Providers(1)

  // match items in content

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
