package main

import (
	"testing"

	"github.com/google/go-github/v86/github"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewClientWithEnv_Nil(t *testing.T) {
	tests := []struct {
		testName string
		funcs    []newClientFunc
	}{
		{
			testName: "empty funcs",
			funcs:    []newClientFunc{},
		},
		{
			testName: "a func returns nil",
			funcs: []newClientFunc{
				func(*zerolog.Logger) *github.Client {
					return nil
				},
			},
		},
		{
			testName: "all funcs return nil",
			funcs: []newClientFunc{
				func(*zerolog.Logger) *github.Client {
					return nil
				},
				func(*zerolog.Logger) *github.Client {
					return nil
				},
			},
		},
	}

	logger := NewNullLogger()

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			c, err := newClientWithEnv(tt.funcs, logger)
			assert.Nilf(t, c, "test_name: %s", tt.testName)
			assert.Error(t, err)
		})
	}
}
