package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"./resources"
	"./dao"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	println("Starting le web service...")
	godotenv.Load();
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	database := createDatabaseConnection()
	defer database.Close()

	postDao := dao.PostDao{Db: database}
	postDao.Init()

	// Registering user resource
	postResource := resources.PostResource{Dao: dao.PostDao{Db: database}}
	postResource.Register(wsContainer)

	println("Listening..")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func createDatabaseConnection() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		println(err.Error)
		panic("failed to connect to database *dies*")
	}
	return db
}