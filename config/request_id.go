package config

import (
	"github.com/google/uuid"
)

type RequestID struct {
	Value string
}

const LOG_KEY_NAME_FOR_REQUEST_ID = "RequestID"

func CreateNewRequestID() RequestID {
	uuid, err := uuid.NewRandom()

	if err != nil {
		return RequestID{"cannot create requestID"}
	}

	return RequestID{uuid.String()}
}
