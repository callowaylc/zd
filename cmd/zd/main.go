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
  app.Logs("application init", app.Entry{
    "config": config,
  })

  retrieved := make(chan struct{ Provider; error }, 15)
  verified := make(chan struct{ Provider; error }, 3)

  // check for new providers and publish to list channel
  go app.Providers(1, list)

  // iterate through providers and check if they already in the system
  for result := range list {
    if result.err != nil {
      app.Logs("error on providers channel", app.Entry{
        "error": result.err,
      })     

    } else {
      go func() {
        if _, err := app.GetProvider(result.Provider.ID); err != nil {
          // provider does not exist; we create it and publish to verified
          app.Logs("failed to get provider", app.Entry{
            "error": err,
            "id": result.Provider.ID,
          })
          verified <- app.CreateProvider(result.Provider, verified)
          
        } else {
          // provider already exists; submit to verified
          verified <- result 
        }
      }()
    }
  }

  for result := range exists {
    if result.err != nil {
      app.Logs("error on exists channel", app.Entry{
        "error": result.err,
      })

    } else {
      go app.CreateProvider(result.Provider)
    }
  }

  // match items in content
  // app.Query("hello")


  os.Exit(0)
}
