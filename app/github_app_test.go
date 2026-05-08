package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClientWithGitHubApp(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal("Failed to generate rsa private key.")
	}

	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	pemDataString := string(pemData)

	tests := []struct {
		testName       string
		appClientID    string
		installationID string
		privateKey     string
		expectNil      bool
	}{
		{
			testName:       "No env at all",
			appClientID:    "",
			installationID: "",
			privateKey:     "",
			expectNil:      true,
		},
		{
			testName:       "Lack of client_id",
			appClientID:    "",
			installationID: "1008",
			privateKey:     pemDataString,
			expectNil:      true,
		},
		{
			testName:       "Lack of installation_id",
			appClientID:    "Iv1.0123456789abcdef",
			installationID: "",
			privateKey:     pemDataString,
			expectNil:      true,
		},
		{
			testName:       "Not integer installation_id",
			appClientID:    "Iv1.0123456789abcdef",
			installationID: "def",
			privateKey:     pemDataString,
			expectNil:      true,
		},
		{
			testName:       "Lack of private_key",
			appClientID:    "Iv1.0123456789abcdef",
			installationID: "1008",
			privateKey:     "",
			expectNil:      true,
		},
		{
			testName:       "Broken private key",
			appClientID:    "Iv1.0123456789abcdef",
			installationID: "1008",
			privateKey:     "foobarbaz",
			expectNil:      true,
		},
		{
			testName:       "Ok",
			appClientID:    "Iv1.0123456789abcdef",
			installationID: "1008",
			privateKey:     pemDataString,
			expectNil:      false,
		},
	}

	logger := NewNullLogger()

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			t.Setenv("GH_APP_CLIENT_ID", tt.appClientID)
			t.Setenv("GH_INSTALLATION_ID", tt.installationID)
			t.Setenv("GH_PRIVATE_KEY", tt.privateKey)

			if tt.expectNil {
				assert.Nil(t, newClientWithGitHubApp(logger))
			} else {
				assert.NotNil(t, newClientWithGitHubApp(logger))
			}
		})
	}
}

func TestNewGitHubAppHTTPClientUsesEmptyPermissions(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	client, err := newGitHubAppHTTPClient("Iv1.0123456789abcdef", 1008, pemData)
	require.NoError(t, err)
	require.NotNil(t, client)

	body, err := json.Marshal(emptyInstallationTokenOptions())
	require.NoError(t, err)
	assert.JSONEq(t, `{"permissions":{}}`, string(body))
}
