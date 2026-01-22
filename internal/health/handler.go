package health

import "github.com/gin-gonic/gin"

// GetHealthStatus handles GET /health requests to check service health, pings the database
func (hc *HealthController) GetHealthStatus(c *gin.Context) {
	err := hc.DB.Ping()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	c.JSON(200, gin.H{"status": "OK"})
}
