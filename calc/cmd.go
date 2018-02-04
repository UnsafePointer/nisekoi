package calc

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type Cmd struct {
	Owner      string
	Repository string
	Username   string
}

func (cmd Cmd) Run() error {
	ctx := context.Background()
	client := github.NewClient(nil)

	repos, error := getRepositories(ctx, client, cmd.Owner)
	if error != nil {
		return error
	}

	fmt.Printf("%v", repos)
	return nil
}

func getRepositories(ctx context.Context, client *github.Client, org string) ([]string, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var apiResponse []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		apiResponse = append(apiResponse, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	var repositories []string
	for _, item := range apiResponse {
		repositories = append(repositories, *item.FullName)
	}

	return repositories, nil
}
