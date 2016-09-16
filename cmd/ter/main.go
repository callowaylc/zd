package main

import (
  "os"
  "runtime"
  "fmt"
  _ "log"
  _ "net/http"
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
  var m = fmt.Sprintf("start application with config %+v", app.GetConfig())
  app.Logs(m, nil)

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
