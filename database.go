package zd

import (
  "fmt"
  "database/sql"
  "os"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseQuery(query string, arguments ...interface{}) (*sql.Rows, error) {
  Logs("query database", Entry{
    "query": query,
    "method": "database.DatabaseQuery",
  })

  statement, err := db.Prepare(query)

  if err != nil {
    Logs("database.DatabaseQuery: failed to prepare query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }
  defer statement.Close()

  rows, err := statement.Query(arguments...)
  if err != nil {
    Logs("database.DatabaseQuery: failed to execute query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }

  Logs("query succeeded", Entry{
    "query": query,
    "method": "database.DatabaseQuery",
  })

  return rows, nil
}

func DatabaseExec(query string, arguments ...interface{}) (sql.Result, error) {
  Logs("query database", Entry{
    "query": query,
    "method": "database.DatabaseExec",
  })

  result, err := db.Exec(query, arguments...)
  if err != nil {
    Logs("database.DatabaseExec: failed to execute query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }

  return result, nil
}

func InitDatabase() error{
  if db == nil {
    config := GetConfig()
    Logs("Opening database connection", Entry{
      "name": config.Name,
      "user": config.User,
      "host": config.Host,
    })
    db, _ = sql.Open("mysql", connectionString())

    if err := db.Ping(); err != nil {
      Logs("Failed to connect to database", Entry{
        "name": config.Name,
        "user": config.User,
        "host": config.Host,
        "error": err,
      })
      os.Exit(1)
    }
  }

  return nil
}

func connectionString() string{
  c := GetConfig()
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.Name)
}
