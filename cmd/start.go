package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/xie392/restful-api/apps/host/http"
	"github.com/xie392/restful-api/apps/host/impl"
	"github.com/xie392/restful-api/conf"
)

var (
	confType string
	confFile string
	confETCD string
)

// StartCmd 表示在没有任何子命令的情况下调用的基本命令
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 host 后端API",
	Long:  "启动 host 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载配置文件
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}

		// 初始化全局日志Logger
		//if err := loadGlobalLogger(); err != nil {
		//	return err
		//}

		// 加载 Host Server 实体
		service := impl.NewHostService()

		// 启动 Host Server
		api := http.NewHostHttpHandler(service)

		// 注册路由
		r := gin.Default()
		api.Registry(r)

		return r.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/config.toml", "Host api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
