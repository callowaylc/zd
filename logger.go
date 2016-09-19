package zd

import (
  "os"
  "time"
  "fmt"
  log "github.com/Sirupsen/logrus"
)

type Entry map[string]interface{}

func init() {
  log.SetFormatter(&log.JSONFormatter{})
  log.SetOutput(os.Stderr)
  log.SetLevel(log.InfoLevel)
}

func Logs (message interface{}, entry Entry) {
  if entry == nil {
    entry = map[string]interface{}{}
  }
  entry["time"] = time.Now()

  log.WithFields(map[string]interface{}(entry)).Info(fmt.Sprintf("%v", message))
}
