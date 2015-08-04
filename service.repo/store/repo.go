package store

type Repo struct {
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	StargazersCount int        `json:"stargazers_count"`
	RepoOwner       *RepoOwner `json:"owner"`
}

type RepoOwner struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
