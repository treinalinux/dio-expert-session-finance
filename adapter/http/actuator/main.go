package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody : Checa a saúde da aplicação
type HealthBody struct {
	Status string `json:"status"`
}

// Health : saúde
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)

}
