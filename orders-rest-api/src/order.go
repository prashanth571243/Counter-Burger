package main

import (
	// "encoding/json"
	// "encoding/json"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"net"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	MongoDBHosts       = "ds143326.mlab.com:43326"
	AuthDatabase       = "cmpe281"
	AuthUserName       = "aditi1203"
	AuthPassword       = "Aditi1203!"
	mongodb_database   = "cmpe281"
	mongodb_collection = "orders"
	// TestDatabase = "goinggo"
)

type BurgerOrder struct {
	OrderId     string `json:"orderId" bson:"orderId"`
	UserId      string `json:"userId" bson:"userId"`
	OrderStatus string `json:"orderStatus" bson:"orderStatus"`
	// Cart        []Items `json:"items" bson:"items"`
	TotalAmount float32 `json:"totalAmount" bson:"totalAmount"`
	// IpAddress   string  `json:"ipaddress" bson:"ipaddress"`
}

type RequiredPayload struct {
	OrderId     string  `json:"orderId" bson:"orderId"`
	UserId      string  `json:"userId" bson:"userId"`
	ItemId      string  `json:"itemId"`
	ItemName    string  `json:"itemName"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

var orders map[string]BurgerOrder

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	server := NewServer()
	server.Run(":" + port)
}

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	router := mux.NewRouter()
	initRoutes(router, formatter)
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	n.UseHandler(handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)(router))
	return n
}

func getIp() string {
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

func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		message := "Burger order API Server Working on machine: "
		//  + getSystemIp()
		formatter.JSON(w, http.StatusOK, struct{ Test string }{message})
	}
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/order/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/orders", getAllBurgers(formatter)).Methods("GET")
	mx.HandleFunc("/order/{orderId}", getBurgerByOrderId(formatter)).Methods("GET")
	mx.HandleFunc("/orders/{userId}", getBurgerByUserId(formatter)).Methods("GET")
	mx.HandleFunc("/order", orderBurger(formatter)).Methods("POST")
	// mx.HandleFunc("/order/{orderId}", burgerOrderPaid(formatter)).Methods("PUT")
	// mx.HandleFunc("/order/{orderId}", burgerItemDelete(formatter)).Methods("DELETE")
	// mx.HandleFunc("/order", burgerOrderDelete(formatter)).Methods("DELETE")
}

func getAllBurgers(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// session, _ := mgo.Dial(mongodb_server)
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs: []string{MongoDBHosts},
			// Timeout:  20 * time.Second,
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}
		session, _ := mgo.DialWithInfo(mongoDBDialInfo)
		// session, err := mgo.DialWithInfo(mongoDBDialInfo)
		// if err != nil {
		// 	log.Fatalf("CreateSession: %s\n", err)
		// } else {
		// 	fmt.Printf("successful")
		// }
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		// err:= session.DB("admin").Login(mongo_user, mongo_pass)
		// if err!=nil{
		// 	formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		// 	return
		// }

		var orders_array []BurgerOrder
		err := session.DB(mongodb_database).C(mongodb_collection).Find(bson.M{}).All(&orders_array)
		fmt.Println(session.DB("cmpe281").C("orders").Find(bson.M{}))
		fmt.Println("Burger Orders:", orders_array)
		fmt.Println("Burger check:", err)
		formatter.JSON(w, http.StatusOK, orders_array)
	}
}

func getBurgerByOrderId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// session, _ := mgo.Dial(mongodb_server)
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}
		session, _ := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		// err:= session.DB("admin").Login(mongo_user, mongo_pass)
		// if err!=nil{
		// 	formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		// 	return
		// }

		params := mux.Vars(req)
		var orderid string = params["orderId"]
		fmt.Println("inside get order by order id")
		fmt.Println("orderID: ", orderid)
		var result BurgerOrder
		err := session.DB(mongodb_database).C(mongodb_collection).Find(bson.M{"orderId": orderid}).One(&result)
		if err != nil {
			fmt.Println("err")
			formatter.JSON(w, http.StatusNotFound, "Requested Order Not Found")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Burger Order: ", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

func getBurgerByUserId(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// session, _ := mgo.Dial(mongodb_server)
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}
		session, _ := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		// err:= session.DB("admin").Login(mongo_user, mongo_pass)
		// if err!=nil{
		// 	formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		// 	return
		// }

		params := mux.Vars(req)
		var userid string = params["userId"]
		fmt.Println("Inside get order by user id")
		fmt.Println("userId: ", userid)
		var result BurgerOrder
		err := session.DB(mongodb_database).C(mongodb_collection).Find(bson.M{"userId": userid}).One(&result)
		if err != nil {
			fmt.Println("err")
			formatter.JSON(w, http.StatusNotFound, "Requested UserId Not Found")
			return
		}
		_ = json.NewDecoder(req.Body).Decode(&result)
		fmt.Println("Burger Order: ", result)
		formatter.JSON(w, http.StatusOK, result)
	}
}

func orderBurger(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Open MongoDB Session
		var orderdetail RequiredPayload
		_ = json.NewDecoder(req.Body).Decode(&orderdetail)
		// session, err := mgo.Dial(mongodb_server)
		// if err:= session.DB("admin").Login(mongo_user, mongo_pass); err != nil {
		// 	formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
		// 	return
		// }
		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{MongoDBHosts},
			Database: AuthDatabase,
			Username: AuthUserName,
			Password: AuthPassword,
		}
		session, _ := mgo.DialWithInfo(mongoDBDialInfo)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var order BurgerOrder

		fmt.Println("Orders: ", "Orders not found")
		order = BurgerOrder{
			OrderId:     orderdetail.OrderId,
			UserId:      orderdetail.UserId,
			OrderStatus: "Placed",
			TotalAmount: orderdetail.Price}
		_ = json.NewDecoder(req.Body).Decode(&order)
		err := c.Insert(order)
		if err != nil {
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		fmt.Println("Orders: ", orders)
		formatter.JSON(w, http.StatusOK, order)
	}
}
