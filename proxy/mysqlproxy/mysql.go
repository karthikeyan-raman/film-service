package mysqlproxy

import (
	"film-service/proxy/logger"
	"sync"
	"time"

	// Register mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var once sync.Once

//GetDal --> Returns mysql dal object
func GetDal() (*sqlx.DB, error) {
	var err error
	once.Do(func() {
		db, err = sqlx.Connect("mysql", "root:sakila@tcp(mysql:3306)/sakila")
		logger.Info("Successfully created mysql instance")
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	})
	if err != nil {
		logger.Error("Failed to initialize db")
	}

	return db, nil
}

// ExecuteNamedQuery --> Returns named query to execute
func ExecuteNamedQuery(query string, arg map[string]interface{}) (*sqlx.Rows, error) {
	query, args, _ := sqlx.Named(query, arg)
	query, args, _ = sqlx.In(query, args...)
	dal, _ := GetDal()
	query = dal.Rebind(query)
	rows, queryError := dal.Queryx(query, args...)
	return rows, queryError
}
