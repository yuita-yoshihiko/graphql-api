package testhelper

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func LoadTestDB(workDir string) *sql.DB {
	if err := godotenv.Load(filepath.Join(workDir, ".env")); err != nil {
		panic(err)
	}
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	str, err := pq.ParseURL(databaseURL)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	boil.SetDB(db)
	boil.DebugMode = true
	return db
}

func LoadFixture(workDir, fixtureDir string) *sql.DB {
	db := LoadTestDB(workDir)
	fixture, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		panic(err)
	}
	if err := fixture.Load(); err != nil {
		panic(err)
	}
	return db
}
