package geolocate

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type GMap struct {
	apiKey     string
	httpClient *http.Client
	baseValues url.Values
	baseURl    *url.URL
}

type GMapOption func(*GMap)

const (
	// Reference "https://maps.googleapis.com/maps/api/geocode/json?sensor=false&"
	BaseURL = "https://maps.googleapis.com/maps/api/geocode/json"
)

var GMapInstance *GMap

// PreGMapParam will perpure the url.Values
// k string is the api key
// It will return a url.Values
func PreGMapParam(k string) url.Values {
	v := make(url.Values, 4)
	v.Set("sensor", "false")
	v.Set("key", k)
	return v
}

// WithGMapKey will set the api key of the Google Map API
// k string is the api key of Google Map API
// It will return an option
func WithGMapKey(k string) GMapOption {
	return func(gMapOption *GMap) {
		gMapOption.apiKey = k
		gMapOption.baseValues = PreGMapParam(k)
	}
}

// WithTimeOut will set timeout of Google Map API
// t time.Duration is the expected timeout
// It will return an option
func WithTimeOut(t time.Duration) GMapOption {
	return func(gMapOption *GMap) {
		gMapOption.httpClient = &http.Client{Timeout: t}
	}
}

// NewDefaultGMapOption will set timeout of Google Map API
// t time.Duration is the expected timeout
// It will return an option
func NewDefaultGMapOption() *GMap {
	baseUri, _ := url.Parse(BaseURL)
	return &GMap{
		httpClient: &http.Client{Timeout: time.Second * 10},
		baseURl:    baseUri,
	}
}

// NewGMapInstance will return a instance of GMap client
// gMapOptions are option which will pass setting of GMap client
// return an GMap client
func NewGMapInstance(gMapOptions ...GMapOption) *GMap {
	if nil == GMapInstance {
		gMapWithDefaultOption := NewDefaultGMapOption()
		for _, gOption := range gMapOptions {
			gOption(gMapWithDefaultOption)
		}
		GMapInstance = gMapWithDefaultOption
	}
	return GMapInstance
}

// SetGMapKey will enable caller to set API key
// It the k is missing a error will be throwed
func (g *GMap) SetGMapKey(k string) error {
	if k == "" {
		return errors.New("MISSING_API_KEY")
	}
	g.apiKey = k
	return nil
}

// GetGeoCode will return the an converted lat lng based on address and region
// address string is the address
// reg string is the region code
// Reference : IANA language region subtag
// https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry
// return a LatLng struct
// If the response of Google is incorrect it will also throw an error ( e.g. address not found, google api down etc )
func (g *GMap) GetGeoCode(addr, reg string) (ll LatLng, err error) {
	v := PrepareGMapAddressWithRegionParams(g.baseValues, addr, reg)
	g.baseURl.RawQuery = ""
	g.baseURl.RawQuery = v.Encode()

	resp, err := g.httpClient.Get(g.baseURl.String())
	if nil != err {
		return ll, err
	}

	dump, _ := httputil.DumpResponse(resp, true)
	log.Print(string(dump))

	var gcr GeoCodeResponse
	err = DecodeGeoCodeResponse(resp, &gcr)
	if nil != err {
		return ll, err
	}
	if gcr.Status != "OK" {
		return ll, errors.New(gcr.Status)
	}

	return GetLatLng(&gcr), nil
}

// PrepareGMapAddressWithRegionParams will assign the address and region code
// v url.Values come from previous NewGMapInstance function
// address string is the address
// reg string is the region code
// return a url.Values
func PrepareGMapAddressWithRegionParams(v url.Values, addr, reg string) url.Values {
	v.Set("address", addr)
	v.Set("region", reg)
	return v
}

// DecodeGeoCodeResponse will decode the response from Google Map
// r *http.Response is the response from Google Map
// t interface{} is the struct you would like to bind
// return an error if the format is not match
func DecodeGeoCodeResponse(r *http.Response, t interface{}) error {
	d := json.NewDecoder(r.Body)
	err := d.Decode(&t)
	defer r.Body.Close()
	if nil != err {
		return err
	}
	return nil
}

// GetLatLng will convert Google Map response to string
// gcr is geo code response from Google Map
// return a LatLng Struct
func GetLatLng(gcr *GeoCodeResponse) (ll LatLng) {
	ll.Lat = gcr.Results[0].Geometry.Location.Lat.String()
	ll.Lng = gcr.Results[0].Geometry.Location.Lng.String()
	return
}
