
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"sync"
	"time"
)

const (
	user = "root"
	password = "root1234"
	host = "localhost"
	port = "3306"
	db = "ticket_service"
)

const (
	maxOpenConnection  = 10
	maxIdleConnection  = 5
	maxConnLifeTime    = time.Hour
)

var UserTicketDB *sql.DB

func initMySQL() error {
	var err error
	err = loadConnection(&UserTicketDB)
	if err != nil {
		return err
	}
	return nil
}

func loadConnection(client **sql.DB) error {
	var wg sync.WaitGroup
	wg.Add(1)
	var err error
	*client, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db))
	if err != nil {
		return err
	}
	(*client).SetMaxOpenConns(maxOpenConnection)
	(*client).SetMaxIdleConns(maxIdleConnection)
	(*client).SetConnMaxLifetime(maxConnLifeTime)
	fmt.Fprintf(os.Stderr, "MySQL Connection established with: %s:%s@tcp(%s:%s)/%s\n", user, password, host, port, db)
	return nil
}