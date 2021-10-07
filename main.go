package main

import (
	"fmt"
	"github.com/joho/godotenv"
	database "github.com/serhatmorkoc/go-realworld-example/database"
	"github.com/serhatmorkoc/go-realworld-example/database/seed"
	"github.com/serhatmorkoc/go-realworld-example/handler/api"
	"github.com/serhatmorkoc/go-realworld-example/store"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if err := godotenv.Load("local.env"); err != nil {
		panic("Error loading .env file")
	}

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port,_ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	sd, _ := strconv.ParseBool(os.Getenv("DB_SEED"))
	logo := os.Getenv("CONSOLE_ICON")
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))

	fmt.Println(logo)
	fmt.Printf("driver: %s\n", driver)
	fmt.Printf("host: %s\n", host)
	fmt.Printf("port: %d\n", port)
	fmt.Printf("database: %s\n", dbName)
	fmt.Printf("username: %s\n", username)
	fmt.Printf("password: %s\n", password)
	fmt.Printf("seed: %t\n", sd)

	fmt.Println("------------------------------------")

	db, err := database.Connect(driver, host,dbName,username,password,port,maxConnections)
	if err != nil {
		panic(err)
	}

	us := store.NewUserStore(db)
	cm := store.NewCommentStore(db)

	if sd {
		seed.Seed(us)
	}

	r := api.New(us,cm)
	h := r.Handler()

	if err := http.ListenAndServe(":3000", h); err != nil {
		panic(err)
	}

}
