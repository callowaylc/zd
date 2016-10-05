package zd

import (
  "fmt"
  "errors"
  "regexp"
  "strings"
  "strconv"
  "sync"
)

type Provider struct {
  Name string
  ID int
}
type ProviderCom struct {
  Value Provider
  Err error
}

const ProviderListNum int = 15

func Providers(page int) <-chan ProviderCom {
  Logs("list.Providers", Entry{ "page": page, })

  config := GetConfig()
  pipe := make(chan ProviderCom, ProviderListNum)

  // setup wait group  
  var wg sync.WaitGroup
  wg.Add(ProviderListNum)

  // get list page
  result, err := Memoize(func() (interface{}, error) {
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
    Logs("Failed to memoize providers list", Entry{
      "error": err, 
    })
    pipe <- ProviderCom{
      Provider{}, err,
    }
  }
  source := result.(string)

  // parse list page
  r := regexp.MustCompile("(?sm)<tr.+?review.+?</tr")
  matches := r.FindAllString(source, -1)

  Logs("matched providers", Entry{ "number": len(matches), })

  for i := 0; i < len(matches); i++ {
    go func(match string) {
      defer wg.Done()
      provider := Provider{}

      r = regexp.MustCompile("(?s)id=(?P<ID>[0-9]+)")
      provider.ID, _ = strconv.Atoi(r.FindStringSubmatch(match)[1])

      r = regexp.MustCompile("(?sm)<a.+?>(.+?)</a")
      provider.Name = strings.TrimSpace(r.FindStringSubmatch(match)[1])

      pipe <- ProviderCom{
        provider, nil,
      }      
    }(matches[i])
  }

  go func() {
    wg.Wait()
    Logs("finished matching providers xxgg", nil)
    pipe <- ProviderCom{
      Provider{}, errors.New("closing pipe xxgg"),
    }
    close(pipe)
  }()

  return pipe
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
