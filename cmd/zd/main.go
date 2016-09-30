package main

import (
  "os"
  "runtime"
  _ "fmt"
  _ "log"
  app "github.com/callowaylc/zd"

)

const NumberProviders int = 15  

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

  retrieved := make(chan app.ProviderCom, NumberProviders)
  verified := make(chan app.ProviderCom, 3)

  // check for new providers and publish to list channel
  go app.Providers(1, retrieved)

  // iterate through providers and check if they already in the system
  for result := range retrieved {
    provider := result.Value

    if result.Err != nil {
      app.Logs("error on providers channel", app.Entry{
        "error": result.Err,
      })     

    } else {
      go func(provider app.Provider) {
        if _, err := app.GetProvider(provider.ID); err != nil {
          // provider does not exist; we create it and publish to verified
          app.Logs("failed to get provider", app.Entry{
            "error": err,
            "id": provider.ID,
          })

          _, err := app.CreateProvider(provider.ID, provider.Name)
          if err != nil {
            app.Logs("failed to create provider", app.Entry{
              "error": err,
              "id": provider.ID,
            })
            return
          }          
        } 

        verified <- app.ProviderCom{ provider, nil, }
      }(provider)
    }
  }

  for i := 0; i < cap(retrieved); i++ {
    select {
    case result := <-verified:
      if result.Err != nil {
        app.Logs("error on verified channel", app.Entry{
          "error": result.Err,
        })

      } else {
        provider := result.Value
        go app.CreateProvider(provider.ID, provider.Name)
      }
    }
  }

  os.Exit(0)
}
