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

// FetchProfile => Returns profile data linked to a specific user
func FetchProfile(w http.ResponseWriter, r *http.Request) {
	profile := &models.ProfileSecData{}
	db.Find(&profile)
	json.NewEncoder(w).Encode(&profile)

}

// UpdateProfile => Updates profile sec data
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	profile := &models.ProfileSecData{}
	db.First(&profile)
	json.NewDecoder(r.Body).Decode(profile)
	db.Save(&profile)
	var resp = map[string]interface{}{"status": true, "message": "Profile created"}
	resp["profile"] = profile
	json.NewEncoder(w).Encode(&profile)
}
