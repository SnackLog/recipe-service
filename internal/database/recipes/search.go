package recipes

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/SnackLog/recipe-service/internal/database/models"
)

func Search(db *sql.DB, username, q string) ([]models.Recipe, error) {
	log.Printf("Values: %s %s", username, q)
	searchQuery := strings.Join(strings.Fields(q), " & ") + ":*"
	log.Printf("Search Query: %s", searchQuery)
	query := `SELECT id, name, unit, created_at, username 
				FROM recipes
				WHERE username = $1 AND to_tsvector('german', name) @@ to_tsquery('german', $2)`

	rows, err := db.Query(query, username, searchQuery)

	if err == sql.ErrNoRows {
		return []models.Recipe{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("Error executing search query: %v", err)
	}
	defer rows.Close()

	recipeList := make([]models.Recipe, 0)

	log.Println("Scanning rows...")
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.Id, &recipe.Name, &recipe.Unit, &recipe.CreatedAt, &recipe.Username); err != nil {
			return nil, fmt.Errorf("Error scanning recipe row: %v", err)
		}
		log.Printf("Found recipe: %+v", recipe)
		recipeList = append(recipeList, recipe)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("Error iterating over recipe rows: %v", rows.Err())
	}

	err = populateRecipeIngredients(recipeList, db)
	if err != nil {
		return nil, fmt.Errorf("error populating recipe ingredients: %v", err)
	}

	return recipeList, nil
}

func populateRecipeIngredients(recipeList []models.Recipe, db *sql.DB) error {
	for i := range recipeList {
		err := populateIngredients(db, recipeList[i].Id, &recipeList[i])
		if err != nil {
			return fmt.Errorf("error populating ingredients for recipe %d: %v", recipeList[i].Id, err)
		}
		err = populateCustomIngredients(db, recipeList[i].Id, &recipeList[i])
		if err != nil {
			return fmt.Errorf("error populating custom ingredients for recipe %d: %v", recipeList[i].Id, err)
		}
		log.Printf("Populated recipe: %+v", recipeList[i])
	}
	return nil
}
