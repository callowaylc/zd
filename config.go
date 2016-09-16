package ter

import (
  "io/ioutil"
  yaml "gopkg.in/yaml.v2"
)

type Config struct {
  Login string
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
