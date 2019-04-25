package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = "52.37.99.107:27017"
var mongodb_database = "location"
var mongodb_collection = "location"
var mongo_user = "cmpe281"
var mongo_pass = "cmpe281"
var adminDatabase = "admin"

func getSystemIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr).String()
	address := strings.Split(localAddr, ":")
	fmt.Println("address: ", address[0])
	return address[0]
}

func NewServerConfiguration() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(mx))
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/location/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/location", addLocationHandler(formatter)).Methods("POST")
	mx.HandleFunc("/locations", getAllLocationsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/location/{locationId}", getLocationByIDHandler(formatter)).Methods("GET")
	mx.HandleFunc("/location/{locationId}", deleteLocationHandler(formatter)).Methods("DELETE")
	mx.HandleFunc("/location/zipcode/{zipcode}", getLocationByZipHandler(formatter)).Methods("GET")
	// mx.handleFunc("/restaurant/{locationId}", updateLocationHandler(formatter)).Methods("PUT")
}

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Locations API Server IP: " + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

func getAllLocationsHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		// params := mux.Vars(req)
		// var zipcode string = params["zipcode"]
		// fmt.Println(zipcode);

		var locationArray []Location
		err = collection.Find(bson.M{}).All(&locationArray)

		if locationArray == nil || len(locationArray) <= 0 {
			formatter.JSON(w, http.StatusNotFound, "No restaurants found")
		} else {
			fmt.Println("Result: ", locationArray)
			formatter.JSON(w, http.StatusOK, locationArray)
		}
	}
}

func addLocationHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		uuidForRestaurant, _ := uuid.NewV4()
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		var location Location
		_ = json.NewDecoder(req.Body).Decode(&location)

		location.LocationId = uuidForRestaurant.String()
		fmt.Println("Locations: ", location)
		err = collection.Insert(location)
		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Error occurred. Cannot add restaurant")
			return
		}
		formatter.JSON(w, http.StatusOK, location)
	}
}

func getLocationByIDHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			log.Fatal(err)
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)
		var locationId string = params["locationId"]
		// fmt.Println("All paramaters:", params)
		fmt.Println("location id is : ", locationId)

		var location Location
		err = collection.Find(bson.M{"locationid": locationId}).One(&location)

		if err != nil {
			formatter.JSON(w, http.StatusNotFound, "Cannot find restaurant")
			return
		} else {
			fmt.Println("Result: ", location)
			// res := json.NewEncoder(w).Encode(res)
			formatter.JSON(w, http.StatusOK, location)
		}
	}
}

func getLocationByZipHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Content-Type", "application/json")

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		session.SetMode(mgo.Monotonic, true)
		collection := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)
		var zipcode string = params["zipcode"]
		fmt.Println(zipcode)

		var res []Location
		err = collection.Find(bson.M{"zipcode": zipcode}).All(&res)

		if res == nil || len(res) <= 0 {
			formatter.JSON(w, http.StatusNotFound, "Cannot find any restaurants for that zipcode")
		} else {
			fmt.Println("Result: ", res)
			formatter.JSON(w, http.StatusOK, res)
		}
	}
}

func deleteLocationHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer session.Close()

		if err := session.DB(adminDatabase).Login(mongo_user, mongo_pass); err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)

		params := mux.Vars(req)

		var result Location
		err = c.Find(bson.M{"locationid": params["locationId"]}).One(&result)
		if err == nil {
			c.Remove(bson.M{"locationid": params["locationId"]})
			formatter.JSON(w, http.StatusOK, result)
		} else {
			formatter.JSON(w, http.StatusNotFound, "Restaurant not found for delete")
		}
	}
}
