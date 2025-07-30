package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Validatable interface {
	Validate() error
}

func WithValidation[T any](next func(http.ResponseWriter, *http.Request, T)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto T

		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			http.Error(w, `{"msg":"invalid JSON"}`, http.StatusBadRequest)
			return
		}

		if v, ok := any(dto).(Validatable); ok {
			if err := v.Validate(); err != nil {
				http.Error(w, fmt.Sprintf(`{"msg":"%s"}`, err.Error()), http.StatusBadRequest)
				return
			}
		}

		next(w, r, dto)
	}
}
