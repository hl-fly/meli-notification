package serviceutil

import (
	"encoding/base64"
	"encoding/json"
)

type TokenBody struct {
	Token    string `json:"token"`
	UUID     string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
}

func UnwrapTokenBody(hashAuthToken string) (body TokenBody, err error) {
	dataBytes, err := base64.RawURLEncoding.DecodeString(hashAuthToken)
	if err != nil {
		return body, err
	}

	err = json.Unmarshal(dataBytes, &body)
	if err != nil {
		return body, err
	}

	return body, nil
}
