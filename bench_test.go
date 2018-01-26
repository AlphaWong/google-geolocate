package geolocate_test

// MAP_KEY=<GOOGLE_MAP_KEY> go test -run=-run=BenchmarkGoogleMapInstance -benchmem -cpuprofile cpu.prof -memprofile mem.prof -bench=.

import (
	"testing"

	geo "github.com/AlphaWong/google-geolocate"
)

func NewGMapInstance() {
	geo.NewGMapInstance("")
}

// BenchmarkNewGMapInstance-4      500000000                3.61 ns/op            0 B/op          0 allocs/op
func BenchmarkNewGMapInstance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGMapInstance()
	}
}
