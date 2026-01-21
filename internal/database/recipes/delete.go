package recipes

import "database/sql"

func DeleteRecipeTx(tx *sql.Tx, recipeID int, username string) error {
	// Delete the recipe from the recipes table, cascading deletes will handle related entries
	query := "DELETE FROM recipes WHERE id = $1 AND username = $2"
	_, err := tx.Exec(query, recipeID, username)
	return err
}

func DeleteRecipe(db *sql.DB, recipeID int, username string) error {
	// Delete the recipe from the recipes table, cascading deletes will handle related entries
	query := "DELETE FROM recipes WHERE id = $1 AND username = $2"
	_, err := db.Exec(query, recipeID, username)
	return err
}
