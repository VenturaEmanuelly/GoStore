package repository

import (
	"database/sql"
	"store/internal/controllers"
)

type database struct {
	db *sql.DB
}

func (d database) QueryRow(query string, args []any, dest ...any) error {
	return d.db.QueryRow(query, args...).Scan(dest...)
}

func (d database) Exec(query string, args ...any) (int64, error) {
	res, err := d.db.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func NewDatabase(db *sql.DB) controllers.RepoSql {
	return database{db: db}
}
