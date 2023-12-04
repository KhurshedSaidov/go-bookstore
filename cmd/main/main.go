package main

import (
	"log"
	"net/http"

	"github.com/KhurshedSaidov/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	routers := mux.NewRouter()
	routes.RegisterBookStoreRoutes(routers)
	http.Handle("/", routers)
	log.Fatal(http.ListenAndServe(":9010", routers))
}
