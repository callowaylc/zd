package zd

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Query(query string) (interface{}, error) {
  if db == nil {
    config := GetConfig()
    Logs("Opening database connection", Entry{
      "name": config.Name,
      "user": config.User,
      "host": config.Host,
    })
    db, _ = sql.Open("mysql", connectionString())
    defer db.Close()

    if err := db.Ping(); err != nil {
      Logs("Failed to connect to database", Entry{
        "name": config.Name,
        "user": config.User,
        "host": config.Host,
        "error": err,
      })
      return nil, err
    }
  }

  return nil, nil
}

func connectionString() string{
  c := GetConfig()
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.Name)
}
