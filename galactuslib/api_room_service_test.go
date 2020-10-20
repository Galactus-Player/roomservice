package galactuslib

import (
	"context"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Galactus-Player/roomservice/roomapi"
)

// TestAddRoom tests creating a server, adding a room, and getting the result.
func TestAddRoom(t *testing.T) {
	// first create a new controller
	api := NewRoomApiController()
	roomrouter := NewRouter(api)
	testserver := httptest.NewServer(roomrouter)
	defer testserver.Close()

	// api configuration
	apiConfig := roomapi.NewConfiguration()
	apiConfig.HTTPClient = testserver.Client()
	testurl, err := url.Parse(testserver.URL)
	if err != nil {
		t.Errorf("error parsing url from test server: %s", err)
		return
	}
	apiConfig.Host = testurl.Host

	// create first client
	apic := roomapi.NewAPIClient(apiConfig)
	background := context.Background()
	retRoom, _, err := apic.RoomApi.AddRoom(background)
	if err != nil {
		t.Errorf("error adding room: %s", err)
		return
	}
	t.Logf("room id: %d, room id str: %s\n", retRoom.Id, retRoom.Code)
}

func BenchmarkAddRoom(b *testing.B) {
	// first create a new controller
	api := NewRoomApiController()
	roomrouter := NewRouter(api)
	testserver := httptest.NewServer(roomrouter)
	defer testserver.Close()

	// api configuration
	apiConfig := roomapi.NewConfiguration()
	apiConfig.HTTPClient = testserver.Client()
	testurl, err := url.Parse(testserver.URL)
	if err != nil {
		b.Errorf("error parsing url from test server: %s", err)
		return
	}
	apiConfig.Host = testurl.Host

	// create first client
	apic := roomapi.NewAPIClient(apiConfig)
	background := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err = apic.RoomApi.AddRoom(background)
		if err != nil {
			b.Errorf("error adding room: %s", err)
			return
		}
	}
}
