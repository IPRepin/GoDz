package req

import (
	"godz/5-order-api-auth/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return nil, err
	}

	if err := Validate(body); err != nil {
		res.JsonResponse(w, map[string]string{"error": err.Error()}, http.StatusUnprocessableEntity)
		return nil, err
	}

	return &body, nil
}
