package zd

import (
  "net/http"
  "io/ioutil"
  "fmt"
)

func HttpGet(uri string) (string, http.Header, error) {
  Logs("http.HttpGet", Entry{ "uri": uri, })

  response, err := http.Get(uri)
  if err != nil {
    Logs("Failed HTTP/GET request", Entry{
      "uri": uri,
    })
    return "", nil, err
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)

  if err != nil {
    Logs("Failed to read response body", Entry{
      "uri": uri,
      "response": response,
    })
  }

  fmt.Println(string(body))

  return string(body), response.Header, nil
}
