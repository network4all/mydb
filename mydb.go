package mydb

import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "github.com/network4all/configuration"
   "fmt"
)

var dbconf configuration.Settings
var DB *sql.DB
var dbdebug bool

func InitConfigSettings(config configuration.Settings) {
   // debug enabled?
   if (dbconf.DBdebug == "1") { dbdebug=true }

   // load database config and create database connection
   var err error
   dbconf = config
   connectionString := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", dbconf.DBuser, dbconf.DBpass, dbconf.DBhost, dbconf.DBname)
   if (dbdebug) { fmt.Printf ("database module: connection string '%s' configured\n", connectionString) }
   DB, err = sql.Open("mysql", connectionString)
   checkErr(err)

   // open a connection (ping test)
   err = DB.Ping()
   if err != nil {
      panic(err.Error())
   }
}

func init () {
   dbdebug = false
   // if (dbdebug) { fmt.Printf ("database module: loaded\n") }
}

func Close () {
   // dispose
   defer DB.Close()
   if (dbdebug) { fmt.Printf ("database module: closed connection\n") }
}

func ExecQuery (myquery string) {
   // execute query on existing connection 
   if (dbdebug) { fmt.Printf ("database module: executing query %s\n", myquery) }
   instcmd, err := DB.Prepare (myquery)
   checkErr(err)

   instcmd.Exec()
}

func checkErr(err error) int {
    if err != nil {
        fmt.Println("database module: error occured: %v", err)
        return 1
    } else {
	return 0
    }
}



