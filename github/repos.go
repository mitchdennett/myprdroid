package github

import (
	"context"
	"github.com/google/go-github/v37/github"
	"github.com/mitchdennett/myprdroid"
	"golang.org/x/oauth2"
)

type RepoService struct{

}

func (r *RepoService) Repos(accessToken string) ([]*myprdroid.Repo, error){
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)

	return toRepos(repos), err
}

func toRepos(gitRepos []*github.Repository) []*myprdroid.Repo {
	repos := make([]*myprdroid.Repo, len(gitRepos))
	for i, v := range gitRepos {
		repos[i] = toRepo(v)
	}
	return repos
}

func toRepo(repo *github.Repository) *myprdroid.Repo {
	return &myprdroid.Repo{
		Name: *repo.Name,
	}
}