package main

import (
	DB "QGTodo/pkg/db"
	"QGTodo/pkg/handlers"
	"QGTodo/pkg/util/goEnv"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

type DBConfig struct {
	host     string
	port     int
	user     string
	password string
	database string
}

func sprintfDBConfig(config DBConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.host, config.port, config.user, config.password, config.database)
}
func main() {
	err := godotenv.Load()
	if err != nil {
		if os.IsNotExist(err) {
			log.Print(err)

		} else {
			log.Fatal(err.Error())
		}
	}

	var dbConfig DBConfig

	dbConfig.host, err = goEnv.StrictGetEnv("QGTODO_PG_HOST")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.port, err = goEnv.StrictGetEnvToI("QGTODO_PG_PORT")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.user, err = goEnv.StrictGetEnv("QGTODO_PG_USER")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.password, err = goEnv.StrictGetEnv("QGTODO_PG_PW")
	if err != nil {
		log.Fatal(err)
	}
	dbConfig.database, err = goEnv.StrictGetEnv("QGTODO_PG_DB")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Launching server...")

	db, err := sql.Open("postgres", sprintfDBConfig(dbConfig))
	queries := DB.New(db)
	//queries = Queries{*queries}
	if err != nil {
		panic(err)
	}

	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := httprouter.New()
	router.POST("/signin", handlers.Signin(queries))
	router.POST("/signup", handlers.Signup(queries))
	router.POST("/task", handlers.AddTasksToUser(queries))
	router.GET("/welcome", handlers.Welcome)
	router.GET("/refresh", handlers.Refresh)
	/*	router.POST("/task", CreateTask)
		router.GET("/task", GetTasks)
		router.GET("/task/:task", GetTask)*/
	log.Fatal(http.ListenAndServe(":8000", router))
}
