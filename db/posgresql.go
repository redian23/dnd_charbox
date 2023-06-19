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
	address = viper.Get("POSTGRES_IP").(string)
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

func insertToTable(tableName string, table Table) {
	db, err := connectToDB()

	_, err = db.Exec("INSERT INTO "+tableName+" (data) VALUES($1)", table.Data)
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
func TestPG() {
	db, err := connectToDB()

	type raceSQL struct {
		RaceAbilID  int    `db:"race_abil_id"`
		AbilityName string `db:"ability_name"`
		RaceName    string `db:"race_name"`
		Description string `db:"description"`
	}
	var table raceSQL

	rows, err := db.Query("SELECT t.* FROM public.race_stat_up t ORDER BY race_id LIMIT 501")
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tab2 Table
	for rows.Next() {
		if err = rows.Scan(&table.RaceAbilID, &table.AbilityName, &table.RaceName, &table.Description); err != nil {
			log.Fatal(err)
		}
		tab2 = Table{ID: "<default>",
			Data: Data{
				"race_ability": table.AbilityName,
				"race_name":    table.RaceName,
				"description":  table.Description,
			}}
		insertToTable("race_abilities_json", tab2)
	}

}
