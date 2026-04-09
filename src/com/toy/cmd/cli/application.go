package cli

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func index() {

	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.Run("127.0.0.1:8282")
}
