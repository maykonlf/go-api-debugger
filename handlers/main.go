package handlers

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io/ioutil"
	"net/http"
	"regexp"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methos", "GET, POST, PUT, DELETE, OPTIONS")

	ctx := appengine.NewContext(r)

	body, _ := ioutil.ReadAll(r.Body)
	re := regexp.MustCompile(`\s*\r?\n\s*`)
	body = re.ReplaceAll(body, []byte{})

	headers, _ := json.Marshal(r.Header)

	log.Debugf(ctx, "headers: %s", headers)
	log.Debugf(ctx, "body: %s", body)

	w.WriteHeader(http.StatusOK)
}