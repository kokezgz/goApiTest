package Controllers

import (
	"net/http"
	"time"

	"../Services"
	"../Utils"
	"github.com/gorilla/mux"
)

type Controller struct {
	mongoService Services.IMongoService
	logger       Utils.ILogger
}

func (c *Controller) StartServer() {
	c.inject()

	r := mux.NewRouter()
	r.HandleFunc("/Restaurants", c.handlerRestaurants).Methods("GET", "POST")
	r.HandleFunc("/Restaurants/{id}", c.handlerRestaurant).Methods("GET", "PUT", "DELETE")

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	c.logger.WriteLog("The Server start in port "+srv.Addr, Utils.Info)
	srv.ListenAndServe()
}

//Injections
func (c *Controller) inject() {
	var injService Services.MongoService
	var injLogger Utils.Logger

	c.mongoService = &injService
	c.logger = &injLogger
}
