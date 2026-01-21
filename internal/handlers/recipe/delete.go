package recipe

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

func (rc *RecipeController) Delete(c *gin.Context) {
	username := c.GetString("username")
	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting recipe ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recipe ID"})
		return
	}

	err = recipes.DeleteRecipe(rc.DB, recipeID, username)
	if err != nil {
		log.Printf("Error deleting recipe: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete recipe"})
		return
	}
	c.Status(http.StatusNoContent)
}
