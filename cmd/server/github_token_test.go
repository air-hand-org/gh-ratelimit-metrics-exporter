package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClientOptsWithToken(t *testing.T) {
	tests := []struct {
		testName       string
		ghTokenEnv     string
		githubTokenEnv string
		expectNil      bool
	}{
		{
			testName:       "No env both",
			ghTokenEnv:     "",
			githubTokenEnv: "",
			expectNil:      true,
		},
		{
			testName:       "GH_TOKEN only",
			ghTokenEnv:     "foo",
			githubTokenEnv: "",
			expectNil:      false,
		},
		{
			testName:       "GITHUB_TOKEN only",
			ghTokenEnv:     "",
			githubTokenEnv: "bar",
			expectNil:      false,
		},
		{
			testName:       "Both envs",
			ghTokenEnv:     "foo",
			githubTokenEnv: "bar",
			expectNil:      false,
		},
	}

	logger := NewNullLogger()

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			t.Setenv("GH_TOKEN", tt.ghTokenEnv)
			t.Setenv("GITHUB_TOKEN", tt.githubTokenEnv)

			opts, _ := newClientOptsWithToken(logger)

			if tt.expectNil {
				assert.Nil(t, opts)
			} else {
				assert.NotNil(t, opts)
			}
		})
	}
}
