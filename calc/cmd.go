package calc

import (
	"context"
	"fmt"
	"strings"
	"sync"

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

type Result struct {
	PullRequests []*github.PullRequest
	Err          error
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

	c := make(chan Result, len(repos))
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(val Repository) {
			if cmd.Debug {
				fmt.Printf("Executing goroutine for value: %s", val.Name)
			}
			pullRequests, err := getPullRequests(ctx, client, val)
			c <- Result{PullRequests: pullRequests, Err: err}
			wg.Done()
		}(repo)
	}
	wg.Wait()
	close(c)

	timeAccumulator := float64(0)
	prAccumulator := int64(0)
	userPrAccumulator := int64(0)

	var cmdErr error
	for result := range c {
		pullRequests, err := result.PullRequests, result.Err
		if err != nil {
			cmdErr = err
			break
		}
		for _, pullRequest := range pullRequests {
			prAccumulator++
			if pullRequest.GetMergedAt().IsZero() {
				continue
			}
			if len(cmd.Username) != 0 {
				if !strings.EqualFold(pullRequest.GetUser().GetLogin(), cmd.Username) {
					continue
				}
				userPrAccumulator++
			}
			delta := pullRequest.GetMergedAt().Sub(pullRequest.GetCreatedAt()).Hours()
			timeAccumulator += delta
			if cmd.Debug {
				fmt.Printf("PR: %s\nCreated at: %v\nMerged at:%v\nDelta in hours: %f\n", pullRequest.GetTitle(), pullRequest.GetCreatedAt(), pullRequest.GetMergedAt(), delta)
			}
		}
	}

	userLanded := ""
	average := timeAccumulator / float64(prAccumulator)
	if len(cmd.Username) != 0 {
		userLanded = fmt.Sprintf("%d out of ", userPrAccumulator)
		average = timeAccumulator / float64(userPrAccumulator)
	}
	fmt.Printf("Average landing PR time is: %.2f hours, for a total of %s%d landed PRs\n", average, userLanded, prAccumulator)

	return cmdErr
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
