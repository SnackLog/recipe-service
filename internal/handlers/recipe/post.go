package recipe

import (
	"log"
	"net/http"

	"github.com/SnackLog/recipe-service/internal/database/models"
	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

// Post handles POST /recipe requests to create a new recipe
func (rc *RecipeController) Post(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	recipe.Username = c.GetString("username")

	recipeId, err := recipes.Insert(rc.DB, &recipe)
	if err != nil {
		log.Printf("Error inserting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create recipe"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"recipe_id": recipeId})

}
