package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	restx "github.com/pantheons-ai/pantheon/rest"
)

func init() {
	pflag.StringP("configPath", "-f", "./etc/config.yaml", "config file od service")
}

func main() {
	pflag.Parse()
	viper.AutomaticEnv()
	viper.BindPFlags(pflag.CommandLine)
	configPath := viper.GetString("configPath")

	var c restx.Config
	conf.MustLoad(configPath, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	svcCtx := restx.NewServiceContext(c)
	restx.RegisterHandlers(server, svcCtx)	

	server.Start()
}