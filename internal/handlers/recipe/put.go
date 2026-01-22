package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/recipe-service/internal/database/models"
	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

// Put handles PUT /recipe/:id requests to update an existing recipe in-place
func (rc *RecipeController) Put(c *gin.Context) {
	var recipe models.Recipe
	username := c.GetString("username")
	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error parsing recipe ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	if err := c.ShouldBindJSON(&recipe); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tx, err := rc.DB.Begin()
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to begin transaction"})
		return
	}
	defer tx.Rollback()

	result, err := recipes.DeleteRecipeTx(tx, recipeId, username)
	if err != nil {
		log.Printf("Error deleting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update recipe"})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update recipe"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	recipe.Username = username
	err = recipes.InsertWithTransactionAt(tx, &recipe, recipeId)
	if err != nil {
		log.Printf("Error inserting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update recipe"})
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to commit transaction"})
		return
	}

	c.Status(http.StatusNoContent)
}
