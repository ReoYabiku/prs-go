package infra

import (
	"context"
	"prs-go/entity"

	"github.com/diegosz/go-graphql-client"
	"github.com/morikuni/failure"
	"golang.org/x/oauth2"
)

type GraphQL struct {
	client *graphql.Client
}

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

type repository struct {
	URL             string
	RepositoryOwner string
}

func (g *GraphQL) ListURL(sq *entity.SearchQuery) ([]*entity.URL, error) {
	// TODO: urlとrepository nameの取得

	// TODO: repository nameによる選別
	return nil, nil
}

func (g *GraphQL) QueryRepos() ([]*repository, error) {
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

	var repos []*repository

	for _, repo := range query.User.PullRequests.Nodes {
		repos = append(repos, &repository{
			URL:             string(repo.Url),
			RepositoryOwner: string(repo.Repository.Owner.Login),
		})
	}

	return repos, nil
}
