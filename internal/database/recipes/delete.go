package recipes

import "database/sql"

// DeleteRecipeTx Deletes a recipe with id recipeID belonging to username as part of transaction tx
func DeleteRecipeTx(tx *sql.Tx, recipeID int, username string) (sql.Result, error) {
	// Delete the recipe from the recipes table, cascading deletes will handle related entries
	query := "DELETE FROM recipes WHERE id = $1 AND username = $2"
	result, err := tx.Exec(query, recipeID, username)
	return result, err
}

// DeleteRecipe Atomicaly deletes a recipe with id recipeID belonging to username
func DeleteRecipe(db *sql.DB, recipeID int, username string) (sql.Result, error) {
	// Delete the recipe from the recipes table, cascading deletes will handle related entries
	query := "DELETE FROM recipes WHERE id = $1 AND username = $2"
	result, err := db.Exec(query, recipeID, username)
	return result, err
}
