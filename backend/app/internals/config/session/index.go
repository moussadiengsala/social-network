package tokenjwt

import (
	"net/http"
	"time"

	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/jwt"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
)

type JwtToken struct{}

func (s *JwtToken) Create(w http.ResponseWriter, r *http.Request, response *lib.Response, user models.GetUser) {
	payload := lib.Payload{
		User:           user,
		ExpirationDate: time.Now().Add(24 * time.Hour),
	}

	j := jwt.JWT{
		Payload: payload,
		Header:  jwt.Header{Alg: "HS256", Typ: "JWT", Addr: r.RemoteAddr},
	}

	// Generate JWT using header, user payload, and secret
	token, err := j.Generate()
	if err != nil {
		lib.ErrorWriter(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Data = map[string]string{
		"jwt": token,
	}
}
