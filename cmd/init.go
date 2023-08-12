package main

import (
	"github.com/spf13/viper"
	"log"
	"pregen/pkg/db"
	"pregen/pkg/races"
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

	races.RacePhotoPath = viper.GetString("photoPath")
	htmlSitePath = viper.GetString("htmlSitePath")
	filePath = viper.GetString("filePath")
	logPath = viper.GetString("logPath")
}

func InitServerPathVars(status bool) {
	if status == true {
		db.InitMongoENV("/etc/pregen.d/")
		InitGinENV("/etc/pregen.d/", "prod.cnf")
	} else {
		db.InitMongoENV("../pkg/db/")
		InitGinENV("../configs/", "test.cnf")
	}
}