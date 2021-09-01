package time

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func GetCurrentTime(w http.ResponseWriter, r *http.Request) {
	current_time := time.Now()
	response := map[string]string{
		"current_time": "",
	}

	tz_param := r.URL.Query().Get("tz")
	if strings.Contains(tz_param, ",") {
		response := map[string]string{}
		tz_array := strings.Split(tz_param, ",")
		for _, tz_element := range tz_array {
			loc, err := time.LoadLocation(tz_element)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				response[tz_element] = "invalid timezone"
			} else {
				response[tz_element] = current_time.In(loc).Format("2006-01-02 15:04:05 -0700 MST")
			}
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		loc, err := time.LoadLocation(tz_param)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("invalid timezone")
		} else {
			response["current_time"] = current_time.In(loc).Format("2006-01-02 15:04:05 -0700 MST")

			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}
