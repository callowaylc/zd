package zd

import (
  "fmt"
)

type Profile struct {

}

func GetProfile(provider Provider) (Profile, error) {
  Logs("profile.GetProfile", Entry{ "id": provider.ID, })
  config := GetConfig()

  // get review page
  // var source string
  content, err := Memoize(func() (interface{}, error) {
    body, _, err := HttpGet(fmt.Sprintf("%s?id=%d", config.Profile, provider.ID))
    if err != nil {
      Logs("Failed to GET provider profile", Entry{
        "url": config.Profile,
        "id": provider.ID,
      })
      return nil, err
    }

    return body, nil
  });

  fmt.Printf("%s%s", content)
  return Profile{}, err
}