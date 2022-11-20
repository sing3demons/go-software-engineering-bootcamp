package sql

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	query        string
	rowsAffected int64
	lastInsertId int64
}

func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertId, nil
}

func (m *mockDB) Exec(query string, args ...any) (sql.Result, error) {
	m.query = query
	return m, nil
}

func TestExecQuery(t *testing.T) {
	mock := &mockDB{
		rowsAffected: 32,
	}

	r, _ := execQuery(mock, "SELECT * FROM sql")

	if mock.query != "SELECT * FROM sql" {
		t.Errorf("should have been call db.Exec with query but it not.")
	}

	if r != 32 {
		t.Errorf("should return row effect %d but it got %d.", 32, r)
	}

}
