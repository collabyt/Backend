package handler

import (
	"encoding/json"
	"net/http"

	"github.com/collabyt/Backend/model"
)

func errorStdTreatment(err error, w http.ResponseWriter, httpCode int) {
	w.WriteHeader(httpCode)
	errRet, _ := json.Marshal(
		model.Error{Description: err.Error()},
	)
	w.Write(errRet)
}
