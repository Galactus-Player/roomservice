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

func main() {
	log.Printf("Server started")

	RoomApiController := galactuslib.NewRoomApiController()

	router := galactuslib.NewRouter(RoomApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
