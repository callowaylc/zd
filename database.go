package zd

import (
  "fmt"
  "github.com/ziutek/mymysql/mysql"
  _ "github.com/ziutek/mymysql/native" // Native engine
)

var db *mysql.Conn

func DatabaseQuery(query string, arguments ...interface{}) (*sql.Rows, error) {
  Logs("query database", Entry{
    "query": query,
    "method": "database.DatabaseQuery",
  })
  if err := initdb(); err != nil {
    Logs("database.DatabaseQuery: failed to init query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }

  statement, err := db.Prepare(query)
  if err != nil {
    Logs("database.DatabaseQuery: failed to prepare query", Entry{
      "query": query,
      "error": err,
    })
    return nil, err
  }
  defer statement.Close()
  Logs("prepared statement succeeded", nil)

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

    address := fmt.Sprintf("%s:%s", config.Host, config.Port)
    db := mysql.New("tcp", "", address, config.User, config.Password, config.Name)

    if err := db.Connect(); err != nil {
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
