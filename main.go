package main

import (
	"database/sql"
	"embed"
	"fmt"

	databaseConfig "github.com/SnackLog/database-config-lib"
	serviceConfig "github.com/SnackLog/service-config-lib"
	"github.com/gin-gonic/gin"

	authLib "github.com/SnackLog/auth-lib"
	"github.com/SnackLog/recipe-service/internal/database"
	"github.com/SnackLog/recipe-service/internal/handlers/recipe"
	"github.com/SnackLog/recipe-service/internal/health"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func main() {
	loadConfigs()
	migrateDatabase()

	db := connectDB()
	defer db.Close()

	setupEndpoints(db)

}

func setupEndpoints(db *sql.DB) {
	engine := gin.Default()

	setupRecipeEndpoints(engine, db)
	setupHealthEndpoints(engine, db)

	engine.Run(":80")

}

func setupHealthEndpoints(engine *gin.Engine, db *sql.DB) {
	hc := health.HealthController{DB: db}
	engine.GET("/health", hc.GetHealthStatus)
}

func setupRecipeEndpoints(engine *gin.Engine, db *sql.DB) {
	rc := recipe.RecipeController{DB: db}
	engine.GET("/recipe/:id", authLib.Authentication, rc.Get)
	engine.POST("/recipe", authLib.Authentication, rc.Post)
	engine.PUT("/recipe/:id", authLib.Authentication, rc.Put)
	engine.DELETE("/recipe/:id", authLib.Authentication, rc.Delete)
}

func connectDB() *sql.DB {
	db, err := database.Connect(databaseConfig.GetDatabaseConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}

func migrateDatabase() {
	err := doMigrations()
	if err != nil {
		panic(fmt.Sprintf("Database migration failed: %v", err))
	}
}

func loadConfigs() {
	err := serviceConfig.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = databaseConfig.LoadConfig()
	if err != nil {
		panic(err)
	}
}

// migrationFiles embeds SQL migration files.
//
//go:embed db/migrations/*.sql
var migrationFiles embed.FS

// doMigrations performs database migrations using embedded SQL files.
func doMigrations() error {
	migrationDriver, err := iofs.New(migrationFiles, "db/migrations")
	if err != nil {
		return fmt.Errorf("Failed to create iofs driver: %v", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		migrationDriver,
		databaseConfig.GetDatabaseConnectionString(),
	)

	if err != nil {
		return fmt.Errorf("Failed to create migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to run migrations: %v", err)
	}

	return nil
}
