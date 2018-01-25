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

func TestPassGetGeoCode(t *testing.T) {
	ll, _ := GetGeoCode(getTestingApiKey(), "HKIVETY", "HK")
	assert.Equal(t, "22.342422", ll.Lat)
	assert.Equal(t, "114.106242", ll.Lng)
}

func TestFailGetGeoCode(t *testing.T) {
	ll, err := GetGeoCode(getTestingApiKey(), "SADADASD@!#!@#SADZXCZXCSACADS!", "HK")
	assert.Equal(t, "", ll.Lat)
	assert.Equal(t, "", ll.Lng)
	assert.Equal(t, "ZERO_RESULTS", err.Error())
}

func TestPassGetGeoCodeViaNewGMapInstance(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	ll, _ := g.GetGeoCode("HKIVETY", "HK")
	assert.Equal(t, "22.342422", ll.Lat)
	assert.Equal(t, "114.106242", ll.Lng)
}

func TestPassGetGeoCodeViaNewGMapInstanceInUS(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	ll, _ := g.GetGeoCode("googleplex", "US")
	assert.Equal(t, "37.422000", ll.Lat)
	assert.Equal(t, "-122.084058", ll.Lng)
}

func TestFailSetGMapKeyViaNewGMapInstance(t *testing.T) {
	t.Parallel()
	g := NewGMapInstance(getTestingApiKey())
	err := g.SetGMapKey("")
	assert.Equal(t, "MISSING_API_KEY", err.Error())
}