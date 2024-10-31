package handler

import (
	"fortress/model"
	"fortress/util"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := model.Envelope{
		Message: "System is running",
		Status:  200,
		Data:    nil,
	}

	util.JsonResponse(w, r, &response)
}
