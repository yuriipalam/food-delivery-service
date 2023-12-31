package utils

import (
	"context"
	"database/sql/driver"
	"fmt"
	"food_delivery/config"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
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
		return 0, fmt.Errorf("id must be integer")
	}

	return value, nil
}

func GetIntSliceByKeyFromMuxVars(key string, r *http.Request) ([]int, error) {
	categoryIDsStrSlice := r.URL.Query()[key]

	var categoryIDs []int

	for _, categoryIDStr := range categoryIDsStrSlice {
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			return nil, fmt.Errorf("id must be integers")
		}

		categoryIDs = append(categoryIDs, categoryID)
	}

	return categoryIDs, nil
}

func IsDefaultValue(value interface{}) bool {
	defaultValue := reflect.Zero(reflect.TypeOf(value)).Interface()
	return reflect.DeepEqual(value, defaultValue)
}

func SendCfgToNextMiddleware(cfg *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "cfg", cfg)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}


func Contains(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil // File exists
	}
	if os.IsNotExist(err) {
		return false, nil // File does not exist
	}
	return false, err // Error occurred while checking file existence
}

func СonvertStrSliceToSqlArr(strSlice []string) string {
	for i, str := range strSlice {
		strSlice[i] = "'" + str + "'"
	}

	return "{" + strings.Join(strSlice, ",") + "}"
}

func СonvertDriverValueToStrSlice(value driver.Value) ([]string, error) {
	bytesValue, ok := value.([]byte)
	if !ok {
		return nil, fmt.Errorf("error converting driver.Value to []bytes")
	}

	return strings.Split(strings.Trim(strings.ReplaceAll(string(bytesValue), "\"", ""), "{}"), ","), nil
}
