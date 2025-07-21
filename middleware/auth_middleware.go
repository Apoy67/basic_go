package middleware

import (
	"basic-restfull-golang/halper"
	"basic-restfull-golang/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA1" == r.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		halper.WriteToResponseBody(w, webResponse)
	}
}
