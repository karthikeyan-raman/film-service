package pingcontroller

import (
	"encoding/json"
	"film-service/utils/response"
	"net/http"
)

//Ping --> Check server is up
func Ping(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

//UpgradedPing --> Check server is up
func UpgradedPing(w http.ResponseWriter, r *http.Request) {
	response.Success(w, 200, map[string]bool{"ok": true})
}
