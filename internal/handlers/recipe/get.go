package recipe

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SnackLog/recipe-service/internal/database/models"
	"github.com/SnackLog/recipe-service/internal/database/recipes"
	"github.com/SnackLog/recipe-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type recipeGetResponse struct {
	Id                int                       `json:"id"`
	Name              string                    `json:"name" binding:"required,min=1,max=100"`
	Unit              string                    `json:"unit" binding:"required,min=1,max=50"`
	Username          string                    `json:"-"`
	CreatedAt         time.Time                 `json:"created_at"`
	Ingredients       []models.Ingredient       `json:"ingredients"`
	CustomIngredients []models.CustomIngredient `json:"custom_ingredients"`
}

func mapToRecipeGetResponse(r models.Recipe) recipeGetResponse {
	return recipeGetResponse{
		Id:                r.Id,
		Name:              r.Name,
		Unit:              r.Unit,
		Username:          r.Username,
		CreatedAt:         r.CreatedAt,
		Ingredients:       r.Ingredients,
		CustomIngredients: r.CustomIngredients,
	}
}
func mapToRecipeGetResponseList(recipes []models.Recipe) []recipeGetResponse {
	responseList := make([]recipeGetResponse, len(recipes))
	for i, r := range recipes {
		responseList[i] = mapToRecipeGetResponse(r)
	}
	return responseList
}

// GetRecipes godoc
// @Summary      Search recipes
// @Description  Search recipes for the authenticated user by query string `q`. Optional. If specified, must be at least 3 characters
// @Tags         recipes
// @Produce      json
// @Param        q   query   string  false  "Search query (minimum 3 characters)"
// @Success      200 {array}  recipeGetResponse
// @Failure      400 {object} handlers.Error
// @Failure      500 {object} handlers.Error
// @Security 	 ApiKeyAuth
// @Router       /recipe [get]
func (rc *RecipeController) Get(c *gin.Context) {
	q := strings.TrimSpace(c.Query("q"))
	username := c.GetString("username")

	if q == "" {
		recipeList, err := recipes.GetLatest(rc.DB, username, 100)
		if err != nil {
			log.Println("Error getting latest recipes:", err)
			c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to get latest recipes"})
			return
		}
		c.JSON(http.StatusOK, mapToRecipeGetResponseList(recipeList))
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

	c.JSON(http.StatusOK, mapToRecipeGetResponseList(recipes))

}
