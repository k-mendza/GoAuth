package controllers

import (
	"GoAuth/api/responses"
	"net/http"
)

func (server *Server) Home(responseWriter http.ResponseWriter, request *http.Request) {
	responses.JSON(responseWriter, http.StatusOK, "API version 1")
}
