package handler

import (
	"log"
	"net/http"
	"os"
	"park/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer lib.DefaultCatch(w)
	os.Setenv("TZ", "Asia/Jakarta")
	header := r.Header
	if header.Get("ai-path") != "" {
		myPath := header.Get("ai-path")
		log.Printf("[%s] /api/%s", r.Method, myPath)
		lib.SendJson(map[string]string{"msg": "Hello world"}, w)
	} else {
		log.Printf("[%s] /api/", r.Method)
		panic("Not Found")
	}
}
