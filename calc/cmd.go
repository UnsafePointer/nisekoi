package calc

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Cmd struct {
	Owner       string
	Repository  string
	Username    string
	AccessToken string
	Debug       bool
}

type Repository struct {
	Owner string
	Name  string
}

func (cmd Cmd) Run() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cmd.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	var repos []Repository
	if len(cmd.Repository) != 0 {
		repos = append(repos, Repository{Owner: cmd.Owner, Name: cmd.Repository})
	} else {
		response, err := getRepositories(ctx, client, cmd.Owner)
		if err != nil {
			return err
		}
		repos = append(repos, response...)
	}

	timeAccumulator := float64(0)
	prAccumulator := 0
	for _, repo := range repos {
		pullRequests, err := getPullRequests(ctx, client, repo)
		if err != nil {
			return err
		}
		for _, pullRequest := range pullRequests {
			delta := pullRequest.GetMergedAt().Sub(pullRequest.GetCreatedAt()).Hours()
			if delta < 0 {
				continue
			}
			timeAccumulator += delta
			prAccumulator++
			if cmd.Debug {
				fmt.Printf("PR: %s\nCreated at: %v\nMerged at:%v\nDelta in hours: %f\n", pullRequest.GetTitle(), pullRequest.GetCreatedAt(), pullRequest.GetMergedAt(), delta)
			}
		}
	}

	fmt.Printf("Average landing PR time is: %.2f hours, for a total of %d landed PRs\n", timeAccumulator/float64(prAccumulator), prAccumulator)

	return nil
}

func getPullRequests(ctx context.Context, client *github.Client, repository Repository) ([]*github.PullRequest, error) {
	opt := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
		State:       "closed",
	}

	var apiResponse []*github.PullRequest
	for {
		pullRequests, resp, err := client.PullRequests.List(ctx, repository.Owner, repository.Name, opt)
		if err != nil {
			return nil, err
		}
		apiResponse = append(apiResponse, pullRequests...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return apiResponse, nil
}

func getRepositories(ctx context.Context, client *github.Client, org string) ([]Repository, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
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

	var repositories []Repository
	for _, item := range apiResponse {
		repositories = append(repositories, Repository{Owner: item.GetOwner().GetLogin(), Name: item.GetName()})
	}

	return repositories, nil
}
