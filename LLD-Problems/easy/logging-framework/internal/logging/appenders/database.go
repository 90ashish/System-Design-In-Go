package appenders

import (
	"database/sql"
	"logger/internal/logging"
	"sync"

	_ "github.com/lib/pq"
)

// DatabaseAppender writes logs into a Postgres table.
type DatabaseAppender struct {
	mu   sync.Mutex
	stmt *sql.Stmt
}

// NewDatabaseAppender connects and prepares the INSERT.
func NewDatabaseAppender(dsn string) (*DatabaseAppender, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// Ensure table exists
	if _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS logs (
            id SERIAL PRIMARY KEY,
            timestamp TIMESTAMP,
            level VARCHAR(10),
            message TEXT
        );
    `); err != nil {
		return nil, err
	}
	stmt, err := db.Prepare(`
        INSERT INTO logs (timestamp, level, message) VALUES ($1, $2, $3);
    `)
	if err != nil {
		return nil, err
	}
	return &DatabaseAppender{stmt: stmt}, nil
}

func (dbApp *DatabaseAppender) Append(m logging.LogMessage) error {
	dbApp.mu.Lock()
	defer dbApp.mu.Unlock()
	_, err := dbApp.stmt.Exec(m.Timestamp, m.Level.String(), m.Message)
	return err
}
