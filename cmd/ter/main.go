package main

import (
  "os"
  "runtime"
  "fmt"
  _ "log"
  yaml "gopkg.in/yaml.v2"
  _ "net/http"
  _ "io/ioutil"
  "github.com/callowaylc/ter/config"
)

func init() {
  if cpu := runtime.NumCPU(); cpu == 1 {
    runtime.GOMAXPROCS(2)
  } else {
    runtime.GOMAXPROCS(cpu)
  }
}

func main() {
  var config config.Config
  yaml.Unmarshal([]byte("login: www.msn.com"), &config)
  fmt.Println("%+v", config)
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