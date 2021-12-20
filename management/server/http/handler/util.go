package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// extractAccountIdFromRequestContext extracts accountId from the request context previously filled by the JWT token (after auth)
func extractAccountIdFromRequestContext(r *http.Request) string {
	//token := r.Context().Value("user").(*jwt.Token)
	//claims := token.Claims.(jwt.MapClaims)
	//
	////actually a user id but for now we have a 1 to 1 mapping.
	//return claims["sub"].(string)
	return "test"
}

//writeJSONObject simply writes object to the HTTP reponse in JSON format
func writeJSONObject(w http.ResponseWriter, obj interface{}) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		http.Error(w, "failed handling request", http.StatusInternalServerError)
		return
	}
}

//Duration is used strictly for JSON requests/responses due to duration marshalling issues
type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
