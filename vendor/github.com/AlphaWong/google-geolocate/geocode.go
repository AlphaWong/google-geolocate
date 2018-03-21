package geolocate

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type GMap struct {
	apiKey     string
	httpClient *http.Client
	baseValues url.Values
	baseURl    *url.URL
}

const (
	// Reference "https://maps.googleapis.com/maps/api/geocode/json?sensor=false&"
	BaseURL = "https://maps.googleapis.com/maps/api/geocode/json"
)

var GMapInstance *GMap

func PreGMapParam(k string) url.Values {
	v := make(url.Values, 4)
	v.Set("sensor", "false")
	v.Set("key", k)
	return v
}

func NewGMapInstance(k string) *GMap {
	if nil == GMapInstance {
		baseUri, _ := url.Parse(BaseURL)
		GMapInstance = &GMap{
			apiKey:     k,
			httpClient: &http.Client{Timeout: time.Second * 10},
			baseValues: PreGMapParam(k),
			baseURl:    baseUri,
		}
	}
	return GMapInstance
}

func (g *GMap) SetGMapKey(k string) error {
	if k == "" {
		return errors.New("MISSING_API_KEY")
	}
	g.apiKey = k
	return nil
}

func (g *GMap) GetGeoCode(addr, reg string) (ll LatLng, err error) {
	v := PrepareGMapAddressWithRegionParams(g.baseValues, addr, reg)
	g.baseURl.RawQuery = ""
	g.baseURl.RawQuery = v.Encode()

	resp, err := g.httpClient.Get(g.baseURl.String())
	if nil != err {
		return ll, err
	}

	var gcr GeoCodeResponse
	err = DecodeGeoCodeResponse(resp, &gcr)
	if nil != err {
		return ll, err
	}

	return GetLatLng(&gcr), nil
}

func PrepareGMapAddressWithRegionParams(v url.Values, addr, reg string) url.Values {
	v.Set("address", addr)
	v.Set("region", reg)
	return v
}

func DecodeGeoCodeResponse(r *http.Response, t interface{}) error {
	d := json.NewDecoder(r.Body)
	err := d.Decode(&t)
	defer r.Body.Close()
	if nil != err {
		return err
	}
	return nil
}

func GetLatLng(gcr *GeoCodeResponse) (ll LatLng) {
	ll.Lat = gcr.Results[0].Geometry.Location.Lat.String()
	ll.Lng = gcr.Results[0].Geometry.Location.Lng.String()
	return
}
