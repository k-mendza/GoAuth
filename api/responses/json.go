package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(responseWriter http.ResponseWriter, statusCode int, data interface{}) {
	responseWriter.WriteHeader(statusCode)
	err := json.NewEncoder(responseWriter).Encode(data)
	if err != nil {
		fmt.Fprintf(responseWriter, "%s", err.Error())
	}
}

func ERROR(responseWriter http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(responseWriter, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(responseWriter, http.StatusBadRequest, nil)
}
