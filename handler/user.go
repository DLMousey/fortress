package handler

import (
	"encoding/json"
	"fortress/model"
	"fortress/repo"
	"fortress/util"
	"github.com/go-playground/validator"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	val := validator.New()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	err = val.Struct(user)
	if err != nil {
		log.Println(err.Error())
		util.JsonResponse(w, r, &model.Envelope{
			Status:  400,
			Message: err.Error(),
			Data:    util.GetErrors(err),
		})
		return
	}

	cxn, err := repo.GetConnection()
	if err != nil {
		util.JsonResponse(w, r, &model.Envelope{
			Status:  500,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	record, err := repo.AddUser(cxn.Db, &user)
	if err != nil {
		log.Println(err.Error())
		util.JsonResponse(w, r, &model.Envelope{
			Status:  500,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	util.JsonResponse(w, r, &model.Envelope{
		Status:  201,
		Message: "User created successfully",
		Data:    model.UserToViewModel(*record),
	})

	return
}
