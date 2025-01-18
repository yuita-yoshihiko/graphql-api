package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB
var txKey = struct{}{}

type (
	DBAdministrator interface {
		GetDao(ctx context.Context) boil.ContextExecutor
		Error(err error) error
	}
	dbutils struct {
		db *sql.DB
	}
)

func NewDBAdministrator(db *sql.DB) DBAdministrator {
	return &dbutils{db: db}
}

func Init() {
	connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	connection += " password=" + os.Getenv("DATABASE_PASSWORD") + " sslmode=" + os.Getenv("DATABASE_SSL_MODE")

	DB, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxLifetime(300 * time.Second)

	boil.SetDB(DB)
}

func DoInTx(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	ctx = context.WithValue(ctx, txKey, tx)
	v, err := f(ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return v, nil
}

func (d *dbutils) GetDao(ctx context.Context) boil.ContextExecutor {
	tx, ok := ctx.Value(txKey).(*sql.Tx)
	if ok {
		return tx
	}
	return d.db
}

func (d *dbutils) Error(err error) error {
	if err == nil {
		return nil
	}
	return err
}
