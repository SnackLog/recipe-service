package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

func (rc *RecipeController) Get(c *gin.Context) {
	idParam := c.Param("id")
	username := c.GetString("username")
	recipeId, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
		return
	}

	recipe, err := recipes.GetById(rc.DB, recipeId)

	if err != nil {
		log.Println("Error retrieving recipe:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recipe"})
		return
	}
	if recipe == nil || recipe.Username != username {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	c.JSON(http.StatusOK, recipe)
}
