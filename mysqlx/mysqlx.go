package mysqlx

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

/*
	sql.Tx 和 sql.DB都是一个struct，这里将它们相同的方法定义成一个interface，相当于这两个struct都实现了这个interface
	方便使用该interface做变量进行传递，同时处理sql.Tx sql.DB。方便将事务与非事务的sql操作封装在一个函数里面。
*/
type SqlMethod interface {
	Prepare(query string) (*sql.Stmt, error)
}

type SqlxMethod interface {
	Prepare(query string) (*sql.Stmt, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}


func SqlDB(datasource string, maxConnNum int, maxIdleConn int, maxLifeTime time.Duration) (*sql.DB, error) {
	sqldb, err := sql.Open("mysql", datasource)
	if err != nil {
		return sqldb, err
	}
	sqldb.SetMaxOpenConns(maxConnNum)
	sqldb.SetMaxIdleConns(maxIdleConn)
	sqldb.SetConnMaxLifetime(maxLifeTime)
	if err = sqldb.Ping(); err != nil {
		return sqldb, err
	}
	return sqldb, err
}

func SqlxDB(datasource string, maxConnNum int, maxIdleConn int, maxLifeTime time.Duration) (*sqlx.DB, error) {
	sqlxDB, err := sqlx.Open("mysql", datasource)
	if err != nil {
		return sqlxDB, err
	}
	sqlxDB.SetMaxOpenConns(maxConnNum)
	sqlxDB.SetMaxIdleConns(maxIdleConn)
	sqlxDB.SetConnMaxLifetime(maxLifeTime)
	if err = sqlxDB.Ping(); err != nil {
		return sqlxDB, err
	}

	return sqlxDB, err
}