// Package db a wrapper of Gorm.
package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// GormDB is gorm.DB of gorm package.
type GormDB = gorm.DB

// Constants definition.
const (
	DriverMySQL = "mysql"
)

// Database interface
type DB interface {
	DB() *GormDB
	StdDB() *sql.DB
	Begin() DB
	RollbackUnlessCommitted()
	Commit()
}

// Connect connects to BD.
func Connect(conf *Config) (DB, error) {
	db, err := gorm.Open(conf.Driver, conf.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.LogMode(conf.LogDebug)

	// Config
	if conf.MaxOpenConnections != 0 {
		db.DB().SetMaxOpenConns(conf.MaxOpenConnections)
	}

	if conf.MaxIdleConnections != 0 {
		db.DB().SetMaxIdleConns(conf.MaxIdleConnections)
	}

	if conf.ConnectionMaxLifetime != 0 {
		db.DB().SetConnMaxLifetime(time.Duration(conf.ConnectionMaxLifetime) * time.Second)
	}

	db.SingularTable(true)
	return &database{db}, nil
}

// DB is wrapper of gorm.DB.
type database struct {
	db *gorm.DB
}

// DB returns current instance of *gorm.DB.
func (_this *database) DB() *GormDB {
	return _this.db
}

// StdDB returns generic instance of *sql.DB
func (_this *database) StdDB() *sql.DB {
	return _this.db.DB()
}

// Begin opens a transaction.
func (_this *database) Begin() DB {
	return &database{_this.db.Begin()}
}

// RollbackUnlessCommitted rollbacks if a transaction not committed.
func (_this *database) RollbackUnlessCommitted() {
	_this.db.RollbackUnlessCommitted()
}

// Commit closes and saves a DB transaction.
func (_this *database) Commit() {
	_this.db.Commit()
}

// Config contains connection info of DB.
type Config struct {
	Username string
	Password string
	Host     string
	Port     int64
	Database string
	Driver   string

	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnectionMaxLifetime int
	LogDebug              bool
}

// ConnectionString returns MySQL connection string
func (_this *Config) ConnectionString() string {
	switch _this.Driver {
	case DriverMySQL:
		return fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8mb4",
			_this.Username,
			_this.Password,
			_this.Host,
			_this.Port,
			_this.Database,
		)
	default:
		return ""
	}
}
