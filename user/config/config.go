package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/user/config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}
