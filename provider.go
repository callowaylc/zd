package zd

import (
  _ "fmt"
  "errors"
)

type Provider struct {
  Name string
  ID int
}

func GetProvider(id int) (Provider, error) {
  Logs("Looking up provider", Entry{
    "id": id,
    "function": "provider.GetProvider",
  })
  provider := Provider{}
  rows, err := DatabaseQuery(`
    SELECT id, name
    FROM provider
    WHERE
      id = ?
  `, id)
  if err != nil {
    Logs("provider.GetProvider: failed to query provider", Entry{
      "id": id,
      "error": err,
    })
    return provider, err
  }
  defer rows.Close()

  if rows.Next() {
    if err := rows.Scan(&provider.ID, &provider.Name); err != nil {
      Logs("provider.GetProvider: failed scan for provider", Entry{
        "id": id,
        "error": err,
      })
      return provider, err
    }
  } else {
    return provider, errors.New("could not find provider")
  }

  return provider, nil
}

func CreateProvider(provider Provider) bool{
  return false
}
