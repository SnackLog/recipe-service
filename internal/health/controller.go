package health

import "database/sql"

// HealthController Contains the database connection for health handlers
type HealthController struct {
	DB *sql.DB
}

