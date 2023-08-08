package infra

import (
	"context"
	"fmt"
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
		&oauth2.Token{AccessToken: fmt.Sprintf("Bearer %s", accessToken)},
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

func (g *GraphQL) QueryRepos() (*repository, error) {
	var query struct {
		User struct {
			PullRequests struct {
				Nodes struct {
					Url        graphql.String
					Repository struct {
						Owner struct {
							login graphql.String
						}
					}
				}
			} `graphql:"pullRequests(last: \"10\", states: \"OPEN\")"`
		} `graphql:"user(login: \"ReoYabiku\")"`
	}

	err := g.client.Query(context.Background(), &query, nil)
	if err != nil {
		return nil, failure.Wrap(err)
	}

	return &repository{
		URL:             string(query.User.PullRequests.Nodes.Url),
		RepositoryOwner: string(query.User.PullRequests.Nodes.Repository.Owner.login),
	}, nil
}
