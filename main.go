package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const yes = `
<!DOCTYPE html><html>
<head><style>
.imtw div {
  margin-top: 25%;
}
.imtw div p {
  text-align: center;
  font-family: serif;
  font-size: xx-large;
}
</style><title>Is Min too warm? Let's find out!</title></head>
<body class="imtw">
  <div><p>Yes.</p></div>
</body></html>
<!-- Hello, Min! -->
`

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprint(w, yes)
	})

	hp := fmt.Sprintf(":%v", port)
	log.Printf("Listening on %v", hp)
	http.ListenAndServe(hp, nil)
}
