package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	host     string
	port     int
	user     string
	password string
	database string
}

func portAtoi(port string) int {
	i, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("Error converting Port")
	}
	return i
}
func sprintfDBConfig(config DBConfig) string{
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.database)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := DBConfig{
		host: 	os.Getenv("QGTODO_HOST"),
		port:	portAtoi(os.Getenv("QGTODO_PORT")),
		user:     os.Getenv("QGTODO_USER"),
		password: os.Getenv("QGTODO_PW"),
		database: os.Getenv("QGTODO_DB"),
	}
	fmt.Println("Launching server...")

	db, err := sql.Open("postgres", sprintfDBConfig(dbConfig))
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
