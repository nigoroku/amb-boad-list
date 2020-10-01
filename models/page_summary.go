package models

type PageSummary struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}
