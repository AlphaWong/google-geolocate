[![Build Status](https://travis-ci.org/AlphaWong/google-geolocate.svg?branch=master)](https://travis-ci.org/AlphaWong/google-geolocate)
[![codecov](https://codecov.io/gh/AlphaWong/google-geolocate/branch/master/graph/badge.svg)](https://codecov.io/gh/AlphaWong/google-geolocate)

# About
Golang client for the Google Maps Geocode API

https://developers.google.com/maps/documentation/geocoding/intro


## Usage
```go
import geo "github.com/Alphawong/google-geolocate"
// Slower
ll := geo.GetGeoCode(getTestingApiKey(), "HKIVETY", "HK")

// Faster
g := geo.NewGMapInstance(getTestingApiKey())
ll, _ := geo.GetGeoCode("HKIVETY", "HK")
```

## Geocode ( Faster )
```go
g := geo.NewGMapInstance(getTestingApiKey())
ll, _ := geo.GetGeoCode("HKIVETY", "HK")
fmt.Println(ll)
// Output: {22.342422 114.106242}
```

## Geocode ( Slower )
```go
ll := GetGeoCode(getTestingApiKey(), "HKIVETY", "HK")
fmt.Println(ll)
// Output: {22.342422 114.106242}
```

## Benchmark
run `MAP_KEY=<GOOGLE_MAP_KEY> go test -run=-run=Benchmark -benchmem -cpuprofile cpu.prof -memprofile mem.prof -bench=.`
```
BenchmarkGoogleMapInstance-4    30000000                40.7 ns/op            16 B/op          1 allocs/op
BenchmarkNewGMapInstance-4      500000000                3.79 ns/op            0 B/op          0 allocs/op

```

### License

Under [MIT](LICENSE)

This project is inspired by "github.com/martinlindhe/google-geolocate"
