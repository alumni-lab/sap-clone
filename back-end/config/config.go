package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var varTypes = [5]string{
	"DBHost",
	"DBName",
	"DBPort",
	"DBUser",
	"DBPassword",
}

// RetrieveDBEnvVariables Finds and assigns
// ALL key-values under config.yml to
// env variables
func RetrieveDBEnvVariables() {
	viper.SetConfigName("config")
	// Could add multiple paths if neccesary
	viper.AddConfigPath("../../..")
	viper.SetConfigType("yml")

	// Maybe be able to use this to change
	// the db live for faster testing.
	// viper.WatchConfig()
	// It also has an event handler for changes
	// viper.OnConfigChange(func)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	dbVars := viper.Sub("database")

	// Allows the program to automatically pick up
	// new config variables under database
	for i := range varTypes {
		os.Setenv(varTypes[i], dbVars.GetString(varTypes[i]))
	}
}
