package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/cfunkhouser/ismintoowarm/thisiswhyimhot"
)

type answer struct {
	Answer string
}

type simpleCache struct {
	sync.RWMutex
	latest thisiswhyimhot.MinTemperatureReport
}

var (
	cache simpleCache
)

func refreshCachePeriodically() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		latest := thisiswhyimhot.Latest()
		cache.Lock()
		cache.latest = latest
		cache.Unlock()
	}
}

var (
	answerTmpl = template.Must(template.New("answer").Parse(`
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
  <div><p>{{ .Answer }}</p></div>
</body></html>
<!-- Hello, Min! -->
`))
)

func answerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	cache.RLock()
	latest := cache.latest
	cache.RUnlock()

	var a answer
	if !latest.Success {
		a.Answer = "Probably"
	} else if latest.Temperature < 23.75 {
		a.Answer = "Yes"
	} else {
		a.Answer = "No"
	}

	err := answerTmpl.Execute(w, a)
	if err != nil {
		log.WithError(err).Print("Failed to execute the template.")
		fmt.Fprint(w, "<p>Whoops, something went wrong.</p>")
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set.")
	}

	cache.latest = thisiswhyimhot.Latest()
	go refreshCachePeriodically()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		fmt.Fprint(w, yes)
	})

	hp := fmt.Sprintf(":%v", port)
	log.Printf("Listening on %v", hp)
	http.ListenAndServe(hp, nil)
}
