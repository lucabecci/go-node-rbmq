package pkg

import (
	"encoding/json"

	"github.com/lucabecci/go-node-rbmq/services/receive-services/internal"
)

//TransformData is a function for transform the message of the broker
func TransformData(msg string, data internal.Message) (string, error) {
	err := json.Unmarshal([]byte(msg), data)
	if err != nil {
		return "", err
	}
	json, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
