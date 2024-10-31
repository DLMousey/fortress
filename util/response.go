package util

import (
	"encoding/json"
	"fortress/model"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, e *model.Envelope) {
	res, err := json.Marshal(e)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			panic(err)
		}
	}

	w.WriteHeader(e.Status)
	_, err = w.Write(res)

	if err != nil {
		panic(err)
	}
}
