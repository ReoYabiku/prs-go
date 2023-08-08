package infra

import (
	"os"
	"prs-go/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	graphQL *GraphQL
)

func TestMain(m *testing.M) {
	endpoint := "https://api.github.com/graphql"
	accessToken := os.Getenv("GITHUB_TOKEN")
	graphQL = NewGraphQL(endpoint, accessToken)

	res := m.Run()

	os.Exit(res)
}

func TestListURL(t *testing.T) {
	tests := []struct {
		name   string
		checks func(*testing.T, error, []*entity.URL)
	}{
		{
			name: "not empty",
			checks: func(t *testing.T, err error, got []*entity.URL) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := graphQL.ListURL()

			tt.checks(t, err, got)
		})
	}
}

func TestQueryRepos(t *testing.T) {
	tests := []struct {
		name   string
		checks func(*testing.T, error, []*pullRequest)
	}{
		{
			name: "not empty",
			checks: func(t *testing.T, err error, got []*pullRequest) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := graphQL.QueryRepos()

			tt.checks(t, err, got)
		})
	}
}
