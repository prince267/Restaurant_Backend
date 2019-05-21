package main

import (
	"MAD_Assignment/dbrepository"
	handlerlib "MAD_Assignment/restapplication/packages/httphandlers"
	"MAD_Assignment/restapplication/restaurantcrudhandler"
	mongoutils "MAD_Assignment/utils"
	logger "log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	// "github.com/priteshgudge/gohttpexamples/sample4/dbrepo/userrepo"
	// handlerlib "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/httphandlers"
	// "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/usercrudhandler"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {

	pingHandler := &handlerlib.PingHandler{}
	dbname := "restaurant"
	mongosession, err := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	if err != nil {
		logger.Fatal("Error in creating mongoSession")
	}
	dbsession := dbrepository.NewMongoRepository(mongosession, dbname)
	handler := restaurantcrudhandler.NewRestaurantCrudHandler(dbsession)
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/restaurant/ping", pingHandler)
	h.Handle("/restaurant/", handler)
	h.Handle("/restaurant/{id}", handler)
	// h.Handle("/restaurant/", handler).Methods("POST")
	// h.Handle("/restaurant/{id}", handler).Methods("PUT")
	// h.Handle("/restaurant/{id}", handler).Methods("DELETE")

	//logger.Fatal(http.ListenAndServe(":8080", h))
	logger.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(h)))

}
