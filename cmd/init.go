package main

import (
	"github.com/spf13/viper"
	"log"
	"pregen/pkg/races"
)

var (
	ProdStatus             bool
	htmlSitePath, filePath string
	logPath                string
	imgsPath               string
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

	races.RacePhotoPath = viper.GetString("racePhotoPath")
	htmlSitePath = viper.GetString("htmlSitePath")
	imgsPath = viper.GetString("imgsPath")
	filePath = viper.GetString("filePath")
	logPath = viper.GetString("logPath")
}

func InitServerPathVars() {
	if ProdStatus == true {
		InitGinENV("/etc/charbox.d/", "prod.cnf")
	} else {
		InitGinENV("./configs/", "test.cnf")
	}
}
