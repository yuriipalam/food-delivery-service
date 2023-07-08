package utils

import (
	"context"
	"fmt"
	"food_delivery/config"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"strconv"
)

func GetIDFromMuxVars(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, fmt.Errorf("id must be integer")
	}

	return id, nil
}

func GetIntValueByKeyFromMuxVars(key string, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars[key])
	if err != nil {
		return 0, fmt.Errorf("%s must be integer", key)
	}

	return value, nil
}

func IsDefaultValue(value interface{}) bool {
	defaultValue := reflect.Zero(reflect.TypeOf(value)).Interface()
	return reflect.DeepEqual(value, defaultValue)
}

func SendCfgToMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "cfg", cfg)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
