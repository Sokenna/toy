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
		r.GET("/api/aliyun", func(c *gin.Context) {
			data := map[string]interface{}{
				"cn-qingdao":            "华北1（青岛）",
				"cn-beijing":            "华北2（北京）",
				"cn-zhangjiakou":        "华北3（张家口）",
				"cn-huhehaote":          "华北5（呼和浩特）",
				"cn-wulanchabu":         "华北6（乌兰察布）",
				"cn-hangzhou":           "华东1（杭州）",
				"cn-shanghai":           "华东2（上海）",
				"cn-shenzhen":           "华南1（深圳）",
				"cn-heyuan":             "华南2（河源）",
				"cn-guangzhou":          "华南3（广州）",
				"cn-chengdu":            "西南1（成都）",
				"cn-hongkong":           "中国（香港）",
				"ap-northeast-1":        "亚太东北 1 (东京)",
				"ap-southeast-1":        "亚太东南 1 (新加坡)",
				"ap-southeast-2":        "亚太东南 2 (悉尼)",
				"ap-southeast-3":        "亚太东南 3 (吉隆坡)",
				"ap-southeast-6":        "菲律宾（马尼拉）",
				"ap-southeast-5":        "亚太东南 5 (雅加达)",
				"ap-south-1":            "亚太南部 1 (孟买)",
				"us-east-1":             "美国东部 1 (弗吉尼亚)",
				"us-west-1":             "美国西部 1 (硅谷)",
				"eu-west-1":             "英国 (伦敦)",
				"me-east-1":             "中东东部 1 (迪拜)",
				"eu-central-1":          "欧洲中部 1 (法兰克福)",
				"cn-nanjing":            "华东5（南京-本地地域）",
				"cn-hangzhou-finance":   "杭州 (金融云)",
				"cn-shanghai-finance-1": "上海 (金融云)",
				"cn-shenzhen-finance-1": "深圳 (金融云)",
				"cn-beijing-finance-1":  "北京（金融云）",
			}
			c.JSONP(http.StatusOK, data)
		})
		r.GET("/api/tencent", func(c *gin.Context) {
			data := map[string]interface{}{

				"ap-guangzhou":     "华南地区(广州)",
				"ap-shanghai":      "华东地区(上海)",
				"ap-nanjing":       "华东地区(南京)",
				"ap-hongkong":      "港澳台地区(中国香港)",
				"na-toronto":       "北美地区(多伦多)",
				"ap-beijing":       "华北地区(北京)",
				"ap-singapore":     "亚太东南(新加坡)",
				"ap-bangkok":       "亚太东南(曼谷)",
				"ap-jakarta":       "亚太东南(雅加达)",
				"na-siliconvalley": "美国西部(硅谷)",
				"ap-chengdu":       "西南地区(成都)",
				"ap-chongqing":     "西南地区(重庆)",
				"eu-frankfurt":     "欧洲地区(法兰克福)",
				"eu-moscow":        "欧洲地区(莫斯科)",
				"ap-seoul":         "亚太东北(首尔)",
				"ap-tokyo":         "亚太东北(东京)",
				"ap-mumbai":        "亚太南部(孟买)",
				"na-ashburn":       "美国东部(弗吉尼亚)",
				"sa-saopaulo":      "南美地区(圣保罗)",
			}
			c.JSONP(http.StatusOK, data)
		})

		r.GET("/api/huawei", func(c *gin.Context) {
			data := map[string]interface{}{
				"af-south-1":     "非洲-约翰内斯堡",
				"cn-north-4":     "华北-北京四",
				"cn-north-1":     "华北-北京一",
				"cn-east-2":      "华东-上海二",
				"cn-east-3":      "华东-上海一",
				"cn-south-1":     "华南-广州",
				"cn-south-2":     "华南-深圳",
				"cn-southwest-2": "西南-贵阳一",
				"ap-southeast-2": "亚太-曼谷",
				"ap-southeast-3": "亚太-新加坡",
				"ap-southeast-1": "中国-香港",
			}
			c.JSONP(http.StatusOK, data)
		})
		r.GET("/api/aws", func(c *gin.Context) {
			data := map[string]interface{}{

				"us-gov-west-1":  "AWS GovCloud (US)",
				"us-gov-east-1":  "AWS GovCloud (US-East)",
				"us-east-1":      "US East (N. Virginia)",
				"us-east-2":      "US East (Ohio)",
				"us-west-1":      "US West (N. California)",
				"us-west-2":      "US West (Oregon)",
				"eu-west-1":      "EU (Ireland)",
				"eu-west-2":      "EU (London)",
				"eu-west-3":      "EU (Paris)",
				"eu-central-1":   "EU (Frankfurt)",
				"eu-north-1":     "EU (Stockholm)",
				"eu-south-1":     "EU (Milan)",
				"ap-east-1":      "Asia Pacific (Hong Kong)",
				"ap-south-1":     "Asia Pacific (Mumbai)",
				"ap-southeast-1": "Asia Pacific (Singapore)",
				"ap-southeast-2": "Asia Pacific (Sydney)",
				"ap-southeast-3": "Asia Pacific (Jakarta)",
				"ap-northeast-1": "Asia Pacific (Tokyo)",
				"ap-northeast-2": "Asia Pacific (Seoul)",
				"ap-northeast-3": "Asia Pacific (Osaka)",
				"sa-east-1":      "South America (Sao Paulo)",
				"cn-north-1":     "China (Beijing)",
				"cn-northwest-1": "China (Ningxia)",
				"ca-central-1":   "Canada (Central)",
				"me-south-1":     "Middle East (Bahrain)",
				"af-south-1":     "Africa (Cape Town)",
				"us-iso-east-1":  "US ISO East",
				"us-isob-east-1": "US ISOB East (Ohio)",
				"us-iso-west-1":  "US ISO West",
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
