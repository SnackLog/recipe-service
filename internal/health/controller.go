package health

import "database/sql"

type HealthController struct {
	DB *sql.DB
}

