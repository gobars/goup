package main

import (
	"fmt"
	"github.com/bingoohuang/statiq/fs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"goup"
	_ "goup/statiq"
	"goup/util"
	"strings"
)

func main() {
	util.InitFlags()
	util.StatiqFS, _ = fs.New()
	logrus.SetLevel(logrus.InfoLevel)
	install := viper.GetString("install")
	if len(install) != 0 {
		switch role := install; role {
		case "kafka-ha":
			goup.InstallKafkaHa(util.Config)
		case "mysql-ha":
			goup.InstallMysqlHa(util.Config)
		case "redis-ha":
			goup.InstallRedisHa(util.Config)
		default:
			fmt.Printf("install role [%s] is not support.\n", role)
			pflag.PrintDefaults()
		}
	} else {
		fmt.Printf("Unknown args %s\n", strings.Join(pflag.Args(), " "))
		pflag.PrintDefaults()
	}
}
