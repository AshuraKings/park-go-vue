package handler

import (
	"net/http"
	"os"
	"park/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	os.Setenv("TZ", "Asia/Jakarta")
	lib.SendJson(map[string]string{"msg": "Hellow world"}, w)
}
