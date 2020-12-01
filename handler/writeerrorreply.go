package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/logger"
	"github.com/collabyt/Backend/model"
)

// WriteErrorReply create an error response to the user and add it to the
// responsewriter
func WriteErrorReply(w http.ResponseWriter, httpCode int) {
	w.WriteHeader(httpCode)
	errRet, _ := json.Marshal(
		model.Error{
			ErrorCode:   httpCode,
			Description: http.StatusText(httpCode),
		},
	)
	logger.Error.Printf("Faulty Request: code %d, message \"%s\"", httpCode, http.StatusText(httpCode))
	w.Write(errRet)

}
