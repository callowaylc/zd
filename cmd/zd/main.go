package main

import (
  "os"
  "runtime"
  "sync"
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

  var wg sync.WaitGroup

  // check for new providers and publish to list channel
  retrieved := app.Providers(1)
  verified := make(chan app.ProviderCom, cap(retrieved))

  // iterate through providers and check if they already in the system
  for result:= range retrieved {
    if err, ok := result.Err.(app.ChannelClosed); ok {
      app.Logs("finished verifying retrieved providers", app.Entry{
        "error": err,
      })
      break
    }

    wg.Add(1)
    go func(result app.ProviderCom) {
      defer wg.Done()

      if result.Err != nil {
        app.Logs("error on retrieved channel", app.Entry{
          "error": result.Err,
        })
        return
      }

      provider := result.Value

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
      } else {
        app.Logs("found provider", app.Entry{
          "id": provider.ID,
        })
      }

      verified <- app.ProviderCom{ provider, nil, }
    }(result)
  }
  wg.Wait()
  close(verified)

  // temporary feed into channel of Packet
  comm := make(chan app.Packet, len(verified))

  for v := range verified {
    comm <- app.Packet{
      v.Value, nil,
    }
  }
  close(comm)

  // collect and update provider attributes
  comm = app.ProviderAttributes{}.Process(comm)

  app.Logs("finished verifying providers", nil)
  os.Exit(0)
}
