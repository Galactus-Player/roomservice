package main

import (
	"context"
	"github.com/Galactus-Player/roomservice/roomapi"
	"log"
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
	log.Printf("[AddRoom]: room id: %d, room id str: %s\n", retRoom.Id, retRoom.Code)

	retRoom, _, err = apic.DefaultApi.GetRoomByCode(background, retRoom.Code)
	if err != nil {
		log.Fatalf("error retrieving room: %s\n", err)
		return
	}
	log.Printf("[GetRoomByCode]: room id: %d, room id str: %s\n", retRoom.Id, retRoom.Code)

}
