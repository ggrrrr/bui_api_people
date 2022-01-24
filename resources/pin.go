package resources

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ggrrrr/bui_api_people/utils/egn"
	"github.com/ggrrrr/bui_lib/api"
	"github.com/ggrrrr/bui_lib/token"
)

type PinReq struct {
	Pin string `json:"pin"`
}

func ParsePin(w http.ResponseWriter, r *http.Request) {
	api.SetResponseHeader(w)
	var err error

	body, _ := ioutil.ReadAll(r.Body)
	var req PinReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Printf("errer %v\n", err)
		api.ResponseError(w, 400, "json", err)
		return
	}
	t, err := api.GetAuthorizationBearer(r)
	log.Printf("ParsePin(%v) %v: token: %v", req.Pin, api.GetUserAgent(r.Context()), t)
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	apiClaims, err := token.Verify(t, r.Context())
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	out := egn.Parse(req.Pin)

	log.Printf("ListRequest(%v): token: %v", api.GetUserAgent(r.Context()), apiClaims)
	api.ResponseOk(w, "ok", &out)
}
