package main

import (
	"log"
	"context"
    "github.com/Galactus-Player/roomservice/roomapi"
)

const roomServiceHost string = "localhost:8080"

func main() {

	// create simple api configuration
	apiConfig := roomapi.NewConfiguration()
	apiConfig.Host = roomServiceHost

	// create first client with background context and call server
	apic := roomapi.NewAPIClient(apiConfig)
	background := context.Background()
	retRoom, _, err := apic.RoomApi.AddRoom(background)
	if err != nil {
		log.Fatalf("error adding room: %s\n", err)
		return
	}
	log.Printf("room id: %d, room id str: %s\n", retRoom.Id, retRoom.Code)
}

