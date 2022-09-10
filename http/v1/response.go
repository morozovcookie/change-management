package v1

import (
	"context"
	"encoding/json"
	"net/http"

	cm "github.com/morozovcookie/change-management"
)

func encodeResponse(_ context.Context, writer http.ResponseWriter, status int, response interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	if response == nil {
		return
	}

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}

//nolint:gochecknoglobals
var mapErrorCodeToHTTPStatusCode = map[cm.ErrorCode]int{
	cm.ErrorCodeOK: http.StatusOK,

	cm.ErrorCodeInvalid: http.StatusBadRequest,

	cm.ErrorCodeInternal: http.StatusInternalServerError,
}

func encodeErrorResponse(ctx context.Context, writer http.ResponseWriter, err error) {
	var (
		code = cm.ErrorCodeFromError(err)
		resp = &struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    code.String(),
			Message: cm.MessageFromError(err),
		}

		status = http.StatusInternalServerError
	)

	if mapped, ok := mapErrorCodeToHTTPStatusCode[code]; ok {
		status = mapped
	}

	encodeResponse(ctx, writer, status, resp)
}
