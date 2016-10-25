package zd

import (
  "io/ioutil"
  "os"
  yaml "gopkg.in/yaml.v2"
)

type Database struct {
  Host string
  Name string
  User string
  Password string
  Port string
}
type Site struct {
  Login string
  List string
  Profile string
}
type Config struct {
  Site
  Database
}
var c Config = Config{}

func InitConfig() {
  contents, err := ioutil.ReadFile("./config.yml")
  if err != nil {
    Logs("Failed to load config", Entry{
      "error": err,
    })
    os.Exit(1)
  }
  yaml.Unmarshal([]byte(contents), c)
}

func GetConfig() Config {
  return c
}
