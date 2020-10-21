/*
 * roomservice
 *
 * An implementation of joinable rooms
 *
 * API version: 0.0.1
 * Contact: decline@umass.edu
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	"github.com/Galactus-Player/roomservice/galactuslib"
)

// TODO(rjected): move this value to a config
const roomServerHost string = "localhost:8080"

func main() {

	RoomApiController := galactuslib.NewRoomApiController()
	router := galactuslib.NewRouter(RoomApiController)

	log.Printf("Server starting on %s\n", roomServerHost)
	log.Fatal(http.ListenAndServe(roomServerHost, router))
}
