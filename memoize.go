package zd

import (
  "fmt"
  "gopkg.in/redis.v4"
)

func Memoize(lambda interface {}) (interface{}, error) {
  client := redis.NewClient(&redis.Options{
      Addr:     "pme-master:9736",
      Password: "", // no password set
      DB:       0,  // use default DB
  })

  _, err := client.Ping().Result()
  if err != nil {
    Logs("Failed to connect to redis", nil)
    return nil, err
  }

  // unpack lambda from interface and generate key
  // based on lambda signature
  f := lambda.(func() (interface{}, error))
  key := fmt.Sprintf("%v", lambda)

  // check if value already exists
  value, err := client.Get(key).Result()

  if err != nil {
    Logs("Failed to retrieve key from redis", Entry{
      "key": key,
    })

    result, _ := f()
    value = fmt.Sprintf("%s", result)

    if err := client.Set(key, value, 0).Err(); err != nil {
      Logs("Failed to memoize value", Entry{
        "key": key,
      })
      return nil, err
    }
  }

  return value, nil
}
