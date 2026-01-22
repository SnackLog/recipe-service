package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

// Delete handles DELETE /recipe/:id requests to delete a recipe by ID
func (rc *RecipeController) Delete(c *gin.Context) {
	username := c.GetString("username")
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting recipe ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe ID"})
		return
	}

	result, err := recipes.DeleteRecipe(rc.DB, recipeID, username)
	if err != nil {
		log.Printf("Error deleting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete recipe"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete recipe"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
