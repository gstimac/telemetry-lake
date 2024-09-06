package models

//github_event : model for a github event

type GithubEvent struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Age         float64 `json:"age"`
	Description string  `json:"description"`
}
