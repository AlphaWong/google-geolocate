package geolocate

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestingApiKey() string {
	key := os.Getenv("MAP_KEY")
	if key == "" {
		log.Println("Missing testing api key")
	}
	return key
}

func TestPassGetGeoCodeViaNewGMapInstance(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	ll, _ := g.GetGeoCode("HKIVETY", "HK")
	assert.Equal(t, "22.342422", ll.Lat)
	assert.Equal(t, "114.1062419", ll.Lng)
}

func TestPassGetGeoCodeViaNewGMapInstanceInUS(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	ll, _ := g.GetGeoCode("googleplex", "US")
	assert.Equal(t, "37.4219999", ll.Lat)
	assert.Equal(t, "-122.0840575", ll.Lng)
}

// This library will always return the first result even the address is ambiguous.
func TestPassGetGeoCodeViaNewGMapInstanceInUSWithAmbiguousAddress(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	ll, _ := g.GetGeoCode("building", "US")
	assert.Equal(t, "37.5292712", ll.Lat)
	assert.Equal(t, "-95.6216764", ll.Lng)
}

func TestFailSetGMapKeyViaNewGMapInstance(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	err := g.SetGMapKey("")
	assert.Equal(t, "MISSING_API_KEY", err.Error())
}
