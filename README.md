[![Build Status](https://travis-ci.org/AlphaWong/google-geolocate.svg?branch=master)](https://travis-ci.org/AlphaWong/google-geolocate)
[![codecov](https://codecov.io/gh/AlphaWong/google-geolocate/branch/master/graph/badge.svg)](https://codecov.io/gh/AlphaWong/google-geolocate)

# About
Golang client for the Google Maps Geocode API

https://developers.google.com/maps/documentation/geocoding/intro


## Usage
```go
import geo "github.com/Alphawong/google-geolocate"

ll := geo.GetGeoCode(getTestingApiKey(), "HKIVETY", "HK")
```

## Geocode
```go
ll := GetGeoCode(getTestingApiKey(), "HKIVETY", "HK")
fmt.Println(ll)
// Output: {22.342422 114.106242}
```

## Benchmark
run `MAP_KEY=<GOOGLE_MAP_KEY> go test -run=-run=BenchmarkGoogleMapInstance -benchmem -cpuprofile cpu.prof -memprofile mem.prof -bench=.`
```
BenchmarkGoogleMapInstance-4    30000000                39.3 ns/op            16 B/op          1 allocs/op
```

### License

Under [MIT](LICENSE)

This project is inspired by "github.com/martinlindhe/google-geolocate"
