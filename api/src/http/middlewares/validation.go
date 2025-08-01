package middlewares

import (
	"api/src/utils"
	"encoding/json"
	"net/http"
)

type Validatable interface {
	Validate() error
}

func ValidationMiddleware[T any](h func(http.ResponseWriter, *http.Request, T)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dto T

		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Msg: "Invalid request data"})
			return
		}

		if v, ok := any(dto).(Validatable); ok {
			if err := v.Validate(); err != nil {
				utils.WriteJSON(w, http.StatusBadRequest, utils.ErrorResponse{Msg: err.Error()})
				return
			}
		}

		h(w, r, dto)
	})
}
