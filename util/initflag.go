package util

import (
	"fmt"
	"github.com/bingoohuang/statiq/fs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "net/http/pprof" // nolint G108
	"os"
	"strings"
)

var Config Settings
var StatiqFS *fs.StatiqFS

func InitFlags() {
	ver := pflag.BoolP("version", "v", false, "show version")
	conf := pflag.StringP("config", "c", "./config.toml", "config file path")
	pflag.StringP("install", "i", "", `
eg:
install mysql-ga
install redis-ga
install kafka-ga
`)

	pflag.Parse()

	args := pflag.Args()
	if len(args) > 0 {
		fmt.Printf("Unknown args %s\n", strings.Join(args, " "))
		pflag.PrintDefaults()
		os.Exit(1)
	}

	if *ver {
		fmt.Printf("Version: 1.4.0\n")
		return
	}

	viper.SetEnvPrefix("GOUP")
	viper.AutomaticEnv()

	_ = viper.BindPFlags(pflag.CommandLine)

	configFile, _ := homedir.Expand(*conf)
	Config = MustLoadConfig(configFile)
}
