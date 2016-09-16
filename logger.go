package ter

import (
  "os"
  "time"
  log "github.com/Sirupsen/logrus"
)

func init() {
  log.SetFormatter(&log.JSONFormatter{})
  log.SetOutput(os.Stderr)
  log.SetLevel(log.InfoLevel)
}

func Logs (message string, entry map[string]interface{}) {
  if entry == nil {
    entry = map[string]interface{}{}
  }
  entry["time"] = time.Now()

  log.WithFields(entry).Info(message)
}
