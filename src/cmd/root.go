package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "toy",
	Short: "toy is a demo data for bi_monitor system",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()
		r.GET("/api", func(c *gin.Context) {
			data := map[string]interface{}{
				"status": "failure",
				"data":   123,
			}
			c.JSONP(http.StatusOK, data)
		})
		r.Run(":8282")
	},
}
var (
	cfgFile string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "application", "", "config file(default is $HOME/application.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "txl", "author name for copyright attribution")

	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "txl77912204@gmail.com")
	viper.SetDefault("license", "apache")
}
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		fmt.Println(home)
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("application")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//Cobra-CLI is its own program that will create your application and add any commands you want.
//It's the easiest way to incorporate Cobra into your application.
