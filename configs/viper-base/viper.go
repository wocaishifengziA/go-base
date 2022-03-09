package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var name = pflag.StringP("name", "n", "", "help message for flagname")
var age = pflag.Int64P("age", "a", 1234, "help message for flagname")

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	viper.Debug()
	fmt.Println("---------------")
	fmt.Println(viper.Get("name"))
}
