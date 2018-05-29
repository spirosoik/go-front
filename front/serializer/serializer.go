package serializer

import (
	"encoding/json"
	"net/http"
)

//Decode deserialization method of HTTP response to Struct
func Decode(v interface{}) func(r *http.Response) error {
	return func(r *http.Response) error {
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(v); err != nil {
			return err
		}
		return nil
	}
}
