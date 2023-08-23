package main

import (
	"github.com/spf13/viper"
	"log"
)

var (
	htmlSitePath, filePath string
	logPath                string
)

func InitGinENV(envPath, fileENV string) {
	// Set the file name of the configurations file
	viper.SetConfigName(fileENV)
	// Set the path to look for the configurations file
	viper.AddConfigPath(envPath)
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	htmlSitePath = viper.GetString("htmlSitePath")
	filePath = viper.GetString("filePath")
	logPath = viper.GetString("logPath")
}

func InitServerPathVars(status bool) {
	if status == true {
		InitGinENV("/etc/pregen.d/", "prod.cnf")
	} else {
		InitGinENV("../configs/", "test.cnf")
	}
}
