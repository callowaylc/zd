package zd

import (
  "fmt"
  "regexp"
  "os"
)

func Providers(page int) ([]Provider, error) {
  Logs("list.Providers", Entry{ "page": page, })
  config := GetConfig()

  // get list page
  var source string
  content, err := Memoize(func() (interface{}, error) {
    body, _, err := HttpGet(fmt.Sprintf("%s?page=%d", config.List, page))
    if err != nil {
      Logs("Failed to GET provider list", Entry{
        "url": config.List,
        "page": page,
      })
      return nil, err
    }

    return body, nil
  });

  if err != nil {
    Logs("Failed to memoize provider list", nil)
    return nil, err
  }

  source = fmt.Sprintf("%s", content)

  // parse list page
  r, _ := regexp.Compile("(?sm)<tr.+?review.+?</tr")
  matches := r.FindAllString(source, -1)

  for _, match := range matches {
    // match id
    r = regexp.Compile("(?s)id=(?P<ID>[0-9]+)")
    name := r.FindStringSubmatch(match)[1]

    // name
    r = regexp.Compile("(?s)id=(?P<ID>[0-9]+)")
    name := r.FindStringSubmatch(match)[1]

    //fmt.Println(r.FindStringSubmatch(r.FindStringSubmatch(match)))
    os.Exit(0)
  }

  return nil, nil
}
