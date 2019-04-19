package main

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
	"os"
	"net/http"
)

var mongodb_server = os.Getenv("AWS_MONGODB")
var mongodb_database = os.Getenv("MONGODB_DBNAME")
var mongodb_collection = os.Getenv("MONGODB_COLLECTION")
var mongodb_username = os.Getenv("MONGODB_USERNAME")
var mongodb_password = os.Getenv("MONGODB_PASSWORD")

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/gumball", gumballHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/gumball", gumballUpdateHandler(formatter)).Methods("PUT")
	// mx.HandleFunc("/order", gumballNewOrderHandler(formatter)).Methods("POST")
	// mx.HandleFunc("/order/{id}", gumballOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/order", gumballOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/orders", gumballProcessOrdersHandler(formatter)).Methods("POST")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"vCloud9.0 Payments API version 1.0 alive!"})
	}
}
