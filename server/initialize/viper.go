package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"frp-platform/global"
)

func initViper() *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		fmt.Printf("./frp-platform -c [config.yaml]")
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("read config file error: %s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("unmarshal config error: %s", err))
	}
	return v
}
