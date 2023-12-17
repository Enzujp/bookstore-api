package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func parseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshall([] byte(body), x); err != nil {
			return
		}
	}
}