package zd

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DatabaseQuery(query string, arguments ...interface{}) (*sql.Rows, error) {
  if err := initdb(); err != nil {
    Logs("database.DatabaseQuery: failed to init query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }

  rows, err := db.Query(query, arguments...)
  if err != nil {
    Logs("database.DatabaseQuery: failed to execute query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }

  return rows, nil
}

func DatabaseExec(query string, arguments ...interface{}) (sql.Result, error) {
  if err := initdb(); err != nil {
    Logs("database.DatabaseExec: failed to init query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }

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

func initdb() error{
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
      return err
    }
  }

  return nil
}

func connectionString() string{
  c := GetConfig()
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.Name)
}