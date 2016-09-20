package zd

import (
  "io/ioutil"
  yaml "gopkg.in/yaml.v2"
)

type Database struct {
  Host string
  Name string
  User string
  Password string
}
type Site struct {
  Login string
  List string
  Review string
}
type Config struct {
  Site
  Database
}
var c *Config

func GetConfig() Config {
  if c == nil {
    c = &Config{}
    contents, err := ioutil.ReadFile("./config.yml")
    check(err)
    yaml.Unmarshal([]byte(contents), c)
  }

  return *c
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
