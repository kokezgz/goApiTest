package Controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../Api"
	"github.com/gorilla/mux"
)

func (c *Controller) handlerRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.AllRestaurants(s)
		restaurants, _ := json.Marshal(response)
		w.Write(restaurants)

	case "POST":
		var restaurant Api.Restaurant
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&restaurant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}

		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.InsertRestaurant(s, restaurant)
		result, _ := json.Marshal(response)
		w.Write(result)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not Allowed"))
	}
}

func (c *Controller) handlerRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	switch r.Method {
	case "GET":
		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.FindRestaurant(s, id)
		restaurant, _ := json.Marshal(response)
		w.Write(restaurant)

	case "PUT":
		var restaurant Api.Restaurant
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&restaurant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}

		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.UpdateRestaurant(s, restaurant)
		result, _ := json.Marshal(response)
		w.Write(result)

	case "DELETE":
		s := c.mongoService.MongoSession()
		response, _ := c.mongoService.DeleteRestaurant(s, id)
		result, _ := json.Marshal(response)
		w.Write(result)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not Allowed"))
	}
}
