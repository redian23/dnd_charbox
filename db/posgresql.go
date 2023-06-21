package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

var (
	user, passwd, address, dbname string
)

func InitPostgresENV(location string) {
	// Set the file name of the configurations file
	viper.SetConfigName(".env")
	// Set the path to look for the configurations file
	if location == "server" {
		viper.AddConfigPath("./root/.env")
		address = "172.20.0.2" //static ip
	} else {
		viper.AddConfigPath("CICD/")
		address = "172.17.0.2"
	}
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	user = viper.Get("POSTGRES_USER").(string)
	passwd = viper.Get("POSTGRES_PASSWORD").(string)
	dbname = viper.Get("POSTGRES_DB").(string)
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://"+user+":"+passwd+"@"+address+":5432/"+dbname+"?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func CheckPostgresDB() {
	db, err := sql.Open("postgres", "postgres://"+user+":"+passwd+"@"+address+":5432/"+dbname+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[DB] [Connect to Postgres]: ok")
}

func insertToTable(tableName string, table interface{}) {
	db, err := connectToDB()

	_, err = db.Exec("INSERT INTO "+tableName+" (data) VALUES($1)", table)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}
func SelectJsonFromRaceTable() {
	db, _ := connectToDB()

	var table RacePostgresTable
	rows, err := db.Query("SELECT * FROM races_json;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&table.ID, &table.Data); err != nil {
			log.Fatal(err)
		}
		fmt.Println(table.Data.RaceAbility)
	}

}
