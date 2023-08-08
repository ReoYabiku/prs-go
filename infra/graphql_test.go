package infra

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	graphQL *GraphQL
)

func TestMain(m *testing.M) {
	endpoint := "https://api.github.com/graphql"
	accessToken := "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"
	graphQL = NewGraphQL(endpoint, accessToken)

	res := m.Run()

	os.Exit(res)
}

func TestQueryRepos(t *testing.T) {
	tests := []struct {
		name   string
		checks func(*testing.T, error, []*repository)
	}{
		{
			name: "not empty",
			checks: func(t *testing.T, err error, got []*repository) {
				assert.NoError(t, err)
				assert.NotEmpty(t, got, "debug")
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
