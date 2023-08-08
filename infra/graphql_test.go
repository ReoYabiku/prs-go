package infra

import "testing"

func TestMain(m *testing.M) {
	// TODO: 事前処理の記載
}

func TestQueryRepos(t *testing.T) {
	tests := []struct {
		name string
	} {
		{
			name: "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
