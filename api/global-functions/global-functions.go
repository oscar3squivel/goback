package functions

import (
	"encoding/json"
	"net/http"
)

//Function that handle Errors
func CheckError(status int64, err error, res *http.ResponseWriter) bool {
	reponse := make(map[string]interface{})
	if err != nil {
		reponse["ok"] = false
		reponse["status"] = status
		reponse["message"] = err.Error()
		json.NewEncoder(*res).Encode(reponse)
		return true
	}
	return false
}
