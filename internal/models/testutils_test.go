package models

import (
	"database/sql"
	"os"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("mysql", "test_web:pass@/test_snippetbox?parseTime=true&multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		dbCloseErr := db.Close()
		if dbCloseErr != nil {
			t.Fatal(err, dbCloseErr)
		}
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		dbCloseErr := db.Close()
		if dbCloseErr != nil {
			t.Fatal(err, dbCloseErr)
		}
		t.Fatal(err)
	}

	t.Cleanup(func() {
		defer func() {
			err := db.Close()
			if err != nil {
				t.Fatal(err)
			}
		}()

		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}
	})

	return db
}
