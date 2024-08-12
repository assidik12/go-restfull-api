package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(request *http.Request, response interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(response)
	PanicError(err)
}

func WriteResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	PanicError(err)
}
