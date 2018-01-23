package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Init config
func Init() {
	// default filename
	name := "config"
	// check for provided filename
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName(name)   // name of config file (without extension)
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetBool by key
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetString by key
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt by key
func GetInt(n string) int {
	return viper.GetInt(n)
}
