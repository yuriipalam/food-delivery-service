package response

type CategoryResponse struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}
