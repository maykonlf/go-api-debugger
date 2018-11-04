package appengine

import (
	"github.com/maykonlf/go-api-debugger/handlers"
	"net/http"
)

func init() {
	http.HandleFunc("/", handlers.MainHandler)
	//appengine.Main()

	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//}
	//
	//log.Printf("Listening on port %s", port)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
