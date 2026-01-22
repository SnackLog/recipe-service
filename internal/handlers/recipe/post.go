package recipe

import (
	"log"
	"net/http"

	"github.com/SnackLog/recipe-service/internal/database/models"
	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

// Post handles POST /recipe requests to create a new recipe
// @Summary Create a recipe
// @Description Creates a new recipe for the authenticated user.
// @Tags recipe
// @Accept json
// @Produce json
// @Param recipe body models.Recipe true "Recipe object"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /recipe [post]
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
