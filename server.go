package main

import (
	"Roam/Roam_backend/graph"
	"Roam/Roam_backend/graph/generated"
	"Roam/Roam_backend/graph/model"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

const debug bool = false

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No config file cound. Please create a .env file.")
		panic(err)
	}
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db := initDB(debug, port)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB(recreate bool, port string) *gorm.DB {
	//Handle our environment stuff
	host, _ := os.LookupEnv("HOST")
	user, _ := os.LookupEnv("USER")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DBNAME")

	//Create our connection string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password)

	//Open up our database connection
	var err error
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//Handle our migrations
	if recreate {
		db.LogMode(true)

		db.DropTableIfExists("friendships")
		db.DropTableIfExists(&model.Post{}, &model.User{}, &model.NewsPost{})
		db.CreateTable(&model.User{}, &model.Post{}, &model.NewsPost{})
		db.Model(&model.Post{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
		db.Table("friendships").AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
		db.Table("friendships").AddForeignKey("friend_id", "users(id)", "CASCADE", "RESTRICT")
	} else {
		db.AutoMigrate(&model.Post{}, &model.User{}, &model.NewsPost{})
	}

	return db
}
