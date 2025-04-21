package main

import (
	"fmt"
	"net/http"

	"internal-health-check-api/src/utils"
)

func main() {
	http.HandleFunc("/{service}", func(w http.ResponseWriter, r *http.Request) {
		s := r.PathValue("service")
		ch := make(chan bool)
		go utils.HealthPingENVKey(fmt.Sprintf("%s_HEALTH_HOST", s), fmt.Sprintf("%s_HEALTH_PORT", s), fmt.Sprintf("%s_HEALTH_TYPE", s))
		rs := <-ch
		if rs {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service Unavailable"))
	})
	http.ListenAndServe(":8080", nil)
}
