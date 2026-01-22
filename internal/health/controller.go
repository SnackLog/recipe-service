package health

import "database/sql"

// HealthController holds the database connection for health handlers
type HealthController struct {
	DB *sql.DB
}

