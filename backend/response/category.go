package response

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}
