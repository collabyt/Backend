package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/model"
)

// WriteErrorReply create an error response to the user and add it to the
// responsewriter
func WriteErrorReply(w http.ResponseWriter, httpCode int) {
	w.WriteHeader(httpCode)
	errRet, _ := json.Marshal(
		model.Error{Description: http.StatusText(httpCode)},
	)
	w.Write(errRet)
}
