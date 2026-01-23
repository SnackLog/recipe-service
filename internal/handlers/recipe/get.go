package recipe

import (
	"log"
	"net/http"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/gin-gonic/gin"
)

// GetRecipes godoc
// @Summary      Search recipes
// @Description  Search recipes for the authenticated user by query string `q`. The query must be at least 3 characters long.
// @Tags         recipes
// @Produce      json
// @Param        q   query   string  true  "Search query (minimum 3 characters)"
// @Success      200 {array}  interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /recipe [get]
func (rc *RecipeController) Get(c *gin.Context) {
	q := c.Query("q")
	username := c.GetString("username")

	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	if len(q) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' must be at least 3 characters long"})
		return
	}

	recipes, err := recipes.Search(rc.DB, username, q)
	if err != nil {
		log.Println("Error searching recipes:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)

}
