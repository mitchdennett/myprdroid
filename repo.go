package myprdroid

type Repo struct {
	Name string
}

type RepoService interface {
	Repos(accessToken string) ([]*Repo, error)
}
