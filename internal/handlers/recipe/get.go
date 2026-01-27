package recipe

import (
	"log"
	"net/http"
	"strings"

	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/SnackLog/recipe-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// GetRecipes godoc
// @Summary      Search recipes
// @Description  Search recipes for the authenticated user by query string `q`. The query must be at least 3 characters long.
// @Tags         recipes
// @Produce      json
// @Param        q   query   string  true  "Search query (minimum 3 characters)"
// @Success      200 {array}  interface{}
// @Failure      400 {object} handlers.Error
// @Failure      500 {object} handlers.Error
// @Router       /recipe [get]
func (rc *RecipeController) Get(c *gin.Context) {
	q := strings.TrimSpace(c.Query("q"))
	username := c.GetString("username")

	if q == "" {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Query parameter 'q' is required"})
		return
	}

	if len(q) < 3 {
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Query parameter 'q' must be at least 3 characters long"})
		return
	}

	recipes, err := recipes.Search(rc.DB, username, q)
	if err != nil {
		log.Println("Error searching recipes:", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to search recipes"})
		return
	}

	c.JSON(http.StatusOK, recipes)

}
