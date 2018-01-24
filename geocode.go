package geolocate

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func PrepareGMapParam(a, k, r string) url.Values {
	u := make(url.Values, 3)
	u.Set("sensor", "false")
	u.Set("address", a)
	u.Set("key", k)
	u.Set("region", r)
	return u
}

const (
	// Reference "https://maps.googleapis.com/maps/api/geocode/json?sensor=false&"
	BaseURL = "https://maps.googleapis.com/maps/api/geocode/json"
)

var (
	ApiKey string

	h = &http.Client{Timeout: time.Second * 10}
)

func SetApiKey(key string) error {
	if key != "" {
		ApiKey = key
		return nil
	}
	return errors.New("MISSING_API_KEY")
}

func SendGeoCodeRequest(a, r string) (*http.Response, error) {
	baseUri, _ := url.Parse(BaseURL)
	v := PrepareGMapParam(a, ApiKey, r)
	baseUri.RawQuery = v.Encode()
	return h.Get(baseUri.String())
}

func DecodeGeoCodeResponse(r *http.Response, t interface{}) error {
	d := json.NewDecoder(r.Body)
	err := d.Decode(&t)
	if nil != err {
		return err
	}
	return nil
}

func GetLatLng(gcp *GeoCodeResponse) (ll LatLng) {
	ll.Lat = FloatToString(gcp.Results[0].Geometry.Location.Lat)
	ll.Lng = FloatToString(gcp.Results[0].Geometry.Location.Lng)
	return
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 6, 64)
}

func GetGeoCode(k, a, r string) (ll LatLng, err error) {
	if err := SetApiKey(k); nil != err {
		return ll, err
	}

	var gcp GeoCodeResponse
	resp, err := SendGeoCodeRequest(a, r)
	if nil != err {
		return ll, err
	}
	DecodeGeoCodeResponse(resp, &gcp)
	if gcp.Status != "OK" {
		return ll, errors.New(gcp.Status)
	}
	ll = GetLatLng(&gcp)
	return
}
