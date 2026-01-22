package health

import "github.com/gin-gonic/gin"

// GetHealthStatus handles GET /health requests to check service health, pings the database
// @Summary Health check
// @Description Checks service health and database connection.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /recipe/health [get]
func (hc *HealthController) GetHealthStatus(c *gin.Context) {
	err := hc.DB.Ping()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}
