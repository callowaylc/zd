package zd

import (
  "fmt"
  "errors"
  "regexp"
  "strings"
  "strconv"
)

type Provider struct {
  Name string
  ID int
}

func Providers(page int) ([]Provider, error) {
  Logs("list.Providers", Entry{ "page": page, })
  config := GetConfig()

  // get list page
  var source string
  content, err := Memoize(func() (interface{}, error) {
    body, _, err := HttpGet(fmt.Sprintf("%s?page=%d", config.List, page))
    if err != nil {
      Logs("Failed to GET providers list", Entry{
        "url": config.List,
        "page": page,
      })
      return nil, err
    }

    return body, nil
  });

  if err != nil {
    Logs("Failed to memoize providers list", nil)
    return nil, err
  }

  source = fmt.Sprintf("%s", content)

  // parse list page
  r, _ := regexp.Compile("(?sm)<tr.+?review.+?</tr")
  matches := r.FindAllString(source, -1)
  providers := make([]Provider, 15)

  for i := 0; i < len(matches); i++ {
    provider := Provider{}

    // match id
    r = regexp.MustCompile("(?s)id=(?P<ID>[0-9]+)")
    provider.ID, _ = strconv.Atoi(r.FindStringSubmatch(matches[i])[1])

    // name
    r = regexp.MustCompile("(?sm)<a.+?>(.+?)</a")
    provider.Name = strings.TrimSpace(r.FindStringSubmatch(matches[i])[1])

    providers[i] = provider
  }

  Logs("list.Providers: parsed provider list", Entry{
    "providers": providers,
  })

  return providers, nil
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

func CreateProvider(id int, name string) (bool, error) {
  Logs("creating provider", Entry{
    "id": id,
    "name": name,
    "method": "provider.CreateProvider",
  })
  _, err := DatabaseExec(`
    INSERT INTO provider(
      id, name
    ) values (
      ?, ?
    )
  `, id, name)

  if err != nil {
    Logs("failed to insert provider", Entry{
      "id": id,
      "name": name,
      "method": "provider.CreateProvider",
      "error": err,
    })

    return false, err
  }

  return true, nil
}
