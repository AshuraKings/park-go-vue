package lib

import (
	"encoding/json"
	"net/http"
)

func SendJson(resp any, w http.ResponseWriter) {
	jsonData, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
}
