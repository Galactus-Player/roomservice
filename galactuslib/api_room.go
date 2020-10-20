/*
 * roomservice
 *
 * An implementation of joinable rooms
 *
 * API version: 0.0.1
 * Contact: decline@umass.edu
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package galactuslib

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

// A RoomApiController binds http requests to an api service and writes the service results to the http response
type RoomApiController struct {
	service        RoomApiServicer
	defaultservice DefaultApiServicer
	mapLock        sync.Mutex
	roomMap        map[string]Room
}

// NewRoomApiController creates a default api controller
func NewRoomApiController() Router {
	currController := &RoomApiController{
		roomMap: make(map[string]Room),
		mapLock: sync.Mutex{},
	}

	currController.service = NewRoomApiService(&currController.roomMap)
	currController.defaultservice = NewDefaultApiService(&currController.roomMap)

	return currController
}

// Routes returns all of the api route for the RoomApiController
func (c *RoomApiController) Routes() Routes {
	return Routes{
		{
			"AddRoom",
			strings.ToUpper("Post"),
			"/v1/room",
			c.AddRoom,
		},
		{
			"GetRoomByCode",
			strings.ToUpper("Get"),
			"/v1/room/{code}",
			c.GetRoomByCode,
		},
	}
}

// AddRoom - Create a new room
func (c *RoomApiController) AddRoom(w http.ResponseWriter, r *http.Request) {
	c.mapLock.Lock()
	result, err := c.service.AddRoom()
	c.mapLock.Unlock()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}

// GetRoomByCode - get a room by its code
func (c *RoomApiController) GetRoomByCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]
	c.mapLock.Lock()
	result, err := c.defaultservice.GetRoomByCode(code)
	c.mapLock.Unlock()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	EncodeJSONResponse(result, nil, w)
}