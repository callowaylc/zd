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

  providers, err := app.Providers(1)
  if err != nil {
    app.Logs("Failed to retrieve providers list", nil)
    os.Exit(1)
  }

  for _, provider := range providers {
    app.GetProfile(provider)
  }

  // match items in content
  app.Query("hello")


  os.Exit(0)
}
