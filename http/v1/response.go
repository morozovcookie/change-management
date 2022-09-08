package v1

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeResponse(_ context.Context, writer http.ResponseWriter, status int, response interface{}) {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json")

	if response == nil {
		return
	}

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}
