package database

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"runtime"
)

func Init(postgresURL string, debugDb bool) *bun.DB {
	Conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(postgresURL)))

	DataBaseInstance := bun.NewDB(Conn, pgdialect.New(), bun.WithDiscardUnknownColumns())

	maxOpenConnections := 4 * runtime.GOMAXPROCS(0)
	DataBaseInstance.SetMaxOpenConns(maxOpenConnections)
	DataBaseInstance.SetMaxIdleConns(maxOpenConnections)

	// debug в консоли всех запросов к БД
	if true == debugDb {
		DataBaseInstance.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
			bundebug.FromEnv(""),
		))
	}

	checkConnection(DataBaseInstance)

	return DataBaseInstance
}

func checkConnection(db *bun.DB) {
	tables := make([]struct {
		TableName string
	}, 0)
	err := db.NewRaw("SELECT table_name FROM information_schema.tables WHERE table_schema='public'").
		Scan(context.Background(), &tables)
	if err != nil {
		log.Fatal("Database check: ", err)
	}
}
