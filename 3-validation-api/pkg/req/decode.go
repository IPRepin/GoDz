package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var request T
	if err := json.NewDecoder(body).Decode(&request); err != nil {
		return request, err
	}
	return request, nil
}
