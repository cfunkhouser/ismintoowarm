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
func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://thisiswhyimhot.herokuapp.com/", 301)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set.")
	}

        http.HandleFunc("/", redirect)
	hp := fmt.Sprintf(":%v", port)
	log.Printf("Listening on %v", hp)
	http.ListenAndServe(hp, nil)
}
