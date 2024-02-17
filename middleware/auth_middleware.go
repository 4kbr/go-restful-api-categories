package middleware

import (
	"4kbr/restful-golang/helper"
	"4kbr/restful-golang/model/web"
	"net/http"
	"os"
)

type AuthMiddleWare struct {
	Handler http.Handler
}

func NewAuthMiddleWare(handler http.Handler) *AuthMiddleWare {
	return &AuthMiddleWare{Handler: handler}
}

func (middleware *AuthMiddleWare) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	secretKey := os.Getenv("AUTH_KEY")
	if secretKey == request.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		//error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}
