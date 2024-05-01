package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"

	_ "github.com/sijms/go-ora"
	"time"
)

/*

# Usage

## Using transaction

	// Begin transaction
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	// Using transaction
	tx.DB().Read( . . . )

	// Commit transaction
	tx.Commit()

## Using without transaction

	db.Read( . . . )

*/

// GormDB is gorm.DB of gorm package.
type GormDB = gorm.DB

// Errors definition.
var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

// Constants definition.
const (
	DriverMySQL  = "mysql"
	DriverOracle = "oracle"
)

// Connect connects to BD.
func Connect(conf *Config) (*DB, error) {
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

	return &DB{db}, nil
}

// DB is wrapper of gorm.DB.
type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	return &DB{db: db}
}

// DB returns current instance of *gorm.DB.
func (_this *DB) DB() *GormDB {
	return _this.db
}

// Begin opens a transaction.
func (_this *DB) Begin() *DB {
	return &DB{_this.db.Begin()}
}

// RollbackUnlessCommitted rollbacks if a transaction not committed.
func (_this *DB) RollbackUnlessCommitted() {
	_this.db.RollbackUnlessCommitted()
}

// Commit closes and saves a DB transaction.
func (_this *DB) Commit() *gorm.DB {
	return _this.db.Commit()
}

// Config contains connection info of DB.
type Config struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     int64
	Database string
	LogDebug bool

	MaxIdleConnections    int
	MaxOpenConnections    int
	ConnectionMaxLifetime int
}

// OracleDbConfig contains connection details for Oracle
type OracleDbConfig struct {
	Driver                string
	Username              string
	Password              string
	Host                  string
	Port                  int64
	ServiceName           string
	ConnectionMaxIdleTime int64
	ConnectionMaxLifeTime int64
	Unsafe                bool
}

// ConnectionString returns MySQL connection string
func (_this *Config) ConnectionString() string {
	switch _this.Driver {
	case DriverMySQL:
		return fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8mb4&loc=Local",
			_this.Username,
			_this.Password,
			_this.Host,
			_this.Port,
			_this.Database,
		)
	}
	return ""
}

func (_this *OracleDbConfig) ConnectionString() string {
	return fmt.Sprintf("oracle://%s:%s@%s:%d/%s",
		_this.Username,
		_this.Password,
		_this.Host,
		_this.Port,
		_this.ServiceName)
}

func ConnectOracle(config *OracleDbConfig) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(config.Driver, config.ConnectionString())
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	if config.ConnectionMaxLifeTime != 0 {
		conn.SetConnMaxLifetime(time.Duration(config.ConnectionMaxLifeTime) * time.Second)
	}

	if config.ConnectionMaxIdleTime != 0 {
		//	conn.SetConnMaxIdleTime(time.Duration(config.ConnectionMaxIdleTime) * time.Second)
	}

	if config.Unsafe {
		// Disable error on missing fields
		conn = conn.Unsafe()
	}

	return conn, nil
}
