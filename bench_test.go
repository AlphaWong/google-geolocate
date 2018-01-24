package geolocate_test

// MAP_KEY=<GOOGLE_MAP_KEY> go test -run=-run=BenchmarkGoogleMapInstance -benchmem -cpuprofile cpu.prof -memprofile mem.prof -bench=.

import (
	"testing"

	geo "github.com/AlphaWong/google-geolocate"
)

func NewGoogleMapInstance() {
	geo.GetGeoCode("", "a", "r")
}

func BenchmarkGoogleMapInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGoogleMapInstance()
	}
}
