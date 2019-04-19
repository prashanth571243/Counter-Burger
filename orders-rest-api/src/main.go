package main

// import (
// 	"os"
// )./

// func main() {

// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "8000"
// 	}

// 	server := NewServer()
// 	server.Run(":" + port)
// }

// func main() {
// 	// We need this object to establish a session to our MongoDB.
// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:    []string{MongoDBHosts},
// 		Timeout:  20 * time.Second,
// 		Database: AuthDatabase,
// 		Username: AuthUserName,
// 		Password: AuthPassword,
// 	}

// 	// Create a session which maintains a pool of socket connections
// 	// to our MongoDB.
// 	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		log.Fatalf("CreateSession: %s\n", err)
// 	} else {
// 		fmt.Printf("successful")
// 	}
// 	mongoSession.SetMode(mgo.Monotonic, true)

// 	r := mux.NewRouter()
// 	http.ListenAndServe(":8090", r)
// }
