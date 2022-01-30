package resources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ggrrrr/bui_api_people/models"
	"github.com/ggrrrr/bui_api_people/utils/egn"
	"github.com/ggrrrr/bui_lib/api"
	"github.com/ggrrrr/bui_lib/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	api.SetResponseHeader(w)
	var err error

	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("data %v", string(body))
	var req models.Person
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("errer %v\n", err)
		api.ResponseError(w, 400, "json", err)
		return
	}

	t, err := api.GetAuthorizationBearer(r)
	log.Printf("Insert(%v): token: %v", api.GetUserAgent(r.Context()), t)
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	apiClaims, err := token.Verify(t, r.Context())
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	if len(req.FullName) < 4 {
		api.ResponseError(w, 400, "full_name", fmt.Errorf("please set proper names"))
		return
	}

	egnOk := *egn.Parse(req.PIN)
	if egnOk.Ok {
		if req.DateOfBirth.Time().Year() > 0 {
			req.DateOfBirth = primitive.NewDateTimeFromTime(egnOk.DateOfBirth)
		}
		if req.Gender != "" {
			req.Gender = egnOk.Gender
		}
	}
	err = req.Insert(r.Context())
	if err != nil {
		log.Printf("Insert: %v, %v", req, err)
		api.ResponseError(w, 500, "insert error", err)
		return
	}

	log.Printf("Insert(%v): token: %v", api.GetUserAgent(r.Context()), apiClaims)
	api.ResponseOk(w, "ok", nil)

}

func Update(w http.ResponseWriter, r *http.Request) {
	api.SetResponseHeader(w)
	var err error

	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("data %v", string(body))
	var req models.Person
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("errer %v\n", err)
		api.ResponseError(w, 400, "json", err)
		return
	}

	t, err := api.GetAuthorizationBearer(r)
	log.Printf("Update(%v): token: %v", api.GetUserAgent(r.Context()), t)
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	apiClaims, err := token.Verify(t, r.Context())
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	asd := primitive.ObjectID{}
	if req.Id == asd {
		api.ResponseError(w, 400, "wrong id", fmt.Errorf("please set proper id"))
		return
	}
	if len(req.FullName) < 4 {
		api.ResponseError(w, 400, "full_name", fmt.Errorf("please set proper names"))
		return
	}

	person, err := models.GetById(r.Context(), req.Id)
	if err != nil {
		api.ResponseError(w, 400, "not found", fmt.Errorf("person %v not found", req.Id))
		return
	}

	set := bson.M{}

	if len(req.FullName) > 3 {
		set["full_name"] = req.FullName
	}
	if len(req.Labels) > 0 {
		set["labels"] = req.Labels
	}
	if len(req.Email) > 0 {
		set["email"] = req.Email
	}
	if len(req.Name) > 0 {
		set["name"] = req.Name
	}
	var egnOk egn.Egn
	if len(req.PIN) > 0 {
		egnOk = *egn.Parse(req.PIN)
		if egnOk.Ok {
			req.DateOfBirth = primitive.NewDateTimeFromTime(egnOk.DateOfBirth)
			req.Gender = egnOk.Gender
		}
		set["pin"] = req.PIN
	}

	if egnOk.Ok && person.DateOfBirth == 0 && req.DateOfBirth == 0 {
		req.DateOfBirth = primitive.NewDateTimeFromTime(egnOk.DateOfBirth)
	}
	if egnOk.Ok && person.Gender == "" && req.Gender == "" {
		req.DateOfBirth = primitive.NewDateTimeFromTime(egnOk.DateOfBirth)
	}

	if req.DateOfBirth != 0 {
		set["dob"] = req.DateOfBirth
	}
	if req.DateOfBirth != 0 {
		set["gender"] = req.Gender
	}
	update := bson.M{
		"$set": set,
	}
	err = models.Update(r.Context(), req.Id, update)
	if err != nil {
		api.ResponseError(w, 500, "error", err)
	}

	log.Printf("Update(%v): token: %v", api.GetUserAgent(r.Context()), apiClaims)
	api.ResponseOk(w, "ok", nil)
}
