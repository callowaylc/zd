package ter

import (
  "os"
  "time"
  log "github.com/Sirupsen/logrus"
)

type Entry map[string]interface{}

func init() {
  log.SetFormatter(&log.JSONFormatter{})
  log.SetOutput(os.Stderr)
  log.SetLevel(log.InfoLevel)
}

func Logs (message string, entry Entry) {
  if entry == nil {
    entry = map[string]interface{}{}
  }
  entry["time"] = time.Now()

  log.WithFields(map[string]interface{}(entry)).Info(message)
}
