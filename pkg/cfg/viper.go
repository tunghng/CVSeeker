package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func SetupConfig() {
	viper.AddConfigPath("./statics")
	env := os.Getenv("ENVIRONMENT")

	fmt.Printf("CURRENT ENV: %s\n", env)

	viper.SetConfigName("config.default") // config.toml
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file not found")
		} else {
			log.Fatalf("Errorf reading conf file %s", err.Error())
		}
	}
}
