package handler

import (
	"fmt"
	"food_delivery/response"
	"food_delivery/utils"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folder, idStr, name := vars["folder"], vars["id"], vars["name"]

	allowedFolders := []string{"suppliers", "categories", "products"}

	if !utils.Contains(allowedFolders, folder) {
		response.SendNotFoundError(w, fmt.Errorf("img not found"))
		return
	}

	imagePath := fmt.Sprintf("./images/%s/%s/%s", folder, idStr, name)

	result, err := utils.FileExists(imagePath)
	if !result || err != nil {
		response.SendNotFoundError(w, fmt.Errorf("img not found"))
		return
	}

	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		response.SendNotFoundError(w, fmt.Errorf("img not found"))
		return
	}

	contentType := http.DetectContentType(imageData)
	w.Header().Set("Content-Type", contentType)

	_, err = w.Write(imageData)
	if err != nil {
		response.SendInternalServerError(w, fmt.Errorf("failed to write response"))
		return
	}
}
