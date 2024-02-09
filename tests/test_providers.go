package tests

import (
	"github.com/c0deltin/duckdb-driver/duckdb"
	"github.com/samber/lo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"onepixel_backend/src/utils/applogger"
)

func ProvideSqliteDB(dbUrl string, config *gorm.Config) *gorm.DB {
	applogger.Warn("Test: Using sqlite db")
	return lo.Must(gorm.Open(sqlite.Open(dbUrl), config))
}

func ProvideDuckDB(dbUrl string, config *gorm.Config) *gorm.DB {
	applogger.Warn("Test: Using duckdb db")
	return lo.Must(gorm.Open(duckdb.Open(dbUrl), config))
}
