[![Build Status](https://travis-ci.org/AlphaWong/google-geolocate.svg?branch=master)](https://travis-ci.org/AlphaWong/google-geolocate)
[![Test Coverage](https://api.codeclimate.com/v1/badges/89879ec88d0567f002c0/test_coverage)](https://codeclimate.com/github/AlphaWong/google-geolocate/test_coverage)

# About
Golang client for the Google Maps Geocode API

https://developers.google.com/maps/documentation/geocoding/intro

## Test
```sh
MAP_KEY=<GOOGLE_MAP_KEY> go test -short .
```

## Usage
```go
import geo "github.com/Alphawong/google-geolocate"

g := geo.NewGMapInstance(WithGMapKey(getTestingApiKey()))
ll, _ := geo.GetGeoCode("HKIVETY", "HK")
```

## Geocode
```go
g := geo.NewGMapInstance(WithGMapKey(getTestingApiKey()))
ll, _ := geo.GetGeoCode("HKIVETY", "HK")
fmt.Println(ll)
// Output: {22.342422 114.106242}
```

## Benchmark
run `MAP_KEY=<GOOGLE_MAP_KEY> go test -run=Benchmark -benchmem -cpuprofile cpu.prof -memprofile mem.prof -bench=.`
```
BenchmarkNewGMapInstance-4      500000000                3.79 ns/op            0 B/op          0 allocs/op
```

### License

Under [MIT](LICENSE)

This project is inspired by "github.com/martinlindhe/google-geolocate"
