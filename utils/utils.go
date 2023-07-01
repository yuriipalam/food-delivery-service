package utils

import (
	"fmt"
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

func IsDefaultValue(value interface{}) bool {
	defaultValue := reflect.Zero(reflect.TypeOf(value)).Interface()
	return reflect.DeepEqual(value, defaultValue)
}
