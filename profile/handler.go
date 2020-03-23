package profile

import (
	"encoding/json"
	"event_backend/models"
	"event_backend/utils"
	"fmt"
	"net/http"
)

var db = utils.ConnectDB()

// PostProfile => inits profile
func PostProfile(w http.ResponseWriter, r *http.Request) {
	profile := &models.ProfileSecData{}

	json.NewDecoder(r.Body).Decode(profile)

	var resp = map[string]interface{}{"status": true, "message": "Profile created"}
	resp["profile"] = profile

	db.Create(&profile)
	fmt.Println(&resp)

	json.NewEncoder(w).Encode(resp)

}

// FetchProfile => Returns profile data
func FetchProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.ProfileSecData
	db.First(&profile)

	json.NewEncoder(w).Encode(&profile)
}
