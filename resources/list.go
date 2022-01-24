package resources

import (
	"log"
	"net/http"

	"github.com/ggrrrr/bui_api_people/models"
	"github.com/ggrrrr/bui_lib/api"
	"github.com/ggrrrr/bui_lib/token"
	"go.mongodb.org/mongo-driver/bson"
)

func ListRequest(w http.ResponseWriter, r *http.Request) {
	api.SetResponseHeader(w)
	var err error
	t, err := api.GetAuthorizationBearer(r)
	log.Printf("ListRequest(%v): token: %v", api.GetUserAgent(r.Context()), t)
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	apiClaims, err := token.Verify(t, r.Context())
	if err != nil {
		api.ResponseErrorUnauthorized(w, err)
		return
	}
	list, err := models.List(r.Context(), bson.D{})
	if err != nil {
		log.Printf("mnongo error: %v", err)
		api.ResponseError(w, 500, "unable to fetch data", err)
		return
	}

	log.Printf("ListRequest(%v): token: %v", api.GetUserAgent(r.Context()), apiClaims)
	api.ResponseOk(w, "ok", list)
}
