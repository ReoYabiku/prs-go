package infra

import (
	"context"
	"prs-go/entity"
	"prs-go/repository"

	"github.com/diegosz/go-graphql-client"
	"github.com/morikuni/failure"
	"golang.org/x/oauth2"
)

type GraphQL struct {
	client *graphql.Client
}

var _ repository.GitHub = &GraphQL{}

func NewGraphQL(endpoint string, accessToken string) *GraphQL {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := graphql.NewClient(endpoint, httpClient)

	return &GraphQL{
		client: client,
	}
}

type pullRequest struct {
	URL             string
	RepositoryOwner string
}

func (g *GraphQL) ListURL() ([]*entity.URL, error) {
	prs, err := g.QueryRepos()
	if err != nil {
		return nil, failure.Wrap(err)
	}

	var urls []*entity.URL

	for _, pr := range prs {
		if pr.RepositoryOwner == "ficilcom" {
			urls = append(urls, (*entity.URL)(&pr.URL))
		}
	}

	return urls, nil
}

func (g *GraphQL) QueryRepos() ([]*pullRequest, error) {
	var query struct {
		User struct {
			PullRequests struct {
				Nodes []struct {
					Url        graphql.String
					Repository struct {
						Owner struct {
							Login graphql.String
						}
					}
				}
			} `graphql:"pullRequests(last: 10, states: OPEN)"`
		} `graphql:"user(login: \"ReoYabiku\")"`
	}

	err := g.client.Query(context.Background(), &query, nil)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	var prs []*pullRequest

	for _, repo := range query.User.PullRequests.Nodes {
		prs = append(prs, &pullRequest{
			URL:             string(repo.Url),
			RepositoryOwner: string(repo.Repository.Owner.Login),
		})
	}

	return prs, nil
}
