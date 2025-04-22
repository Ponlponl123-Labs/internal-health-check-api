package main

import (
	"fmt"
	"net/http"
	"strings"

	"internal-health-check-api/src/utils"
)

func main() {
	http.HandleFunc("/{service}", func(w http.ResponseWriter, r *http.Request) {
		s := strings.ToUpper(r.PathValue("service"))
		rs := utils.HealthPingENVKey(fmt.Sprintf("%s_HEALTH_HOST", s), fmt.Sprintf("%s_HEALTH_PORT", s), fmt.Sprintf("%s_HEALTH_TYPE", s))
		if !rs {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Service Unavailable"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8080", nil)
}
