package config

import "testing"

func TestBaseUrl(t *testing.T) {
	c := Config{Environment: Sandbox}
	baseURL := c.GetBaseURL()
	if baseURL != "https://connect.squareupsandbox.com" {
		t.Errorf("Unexpected base URL. Expecting %s, Got %s", "https://connect.squareupsandbox.com", baseURL)
	}

	c = Config{Environment: Production}
	baseURL = c.GetBaseURL()
	if baseURL != "https://connect.squareup.com" {
		t.Errorf("Unexpected base URL. Expecting %s, Got %s", "https://connect.squareup.com", baseURL)
	}
}

func TestGetAccessToken(t *testing.T) {
	c := Config{Environment: Sandbox, AccessToken: "1234", SandboxAccessToken: "ERROR"}
	accessToken, _ := c.GetAccessToken()
	if accessToken != "1234" {
		t.Errorf("Unexpected Access Token. Expecting 1234, Got %s", accessToken)
	}

	c = Config{Environment: Sandbox, SandboxAccessToken: "1234"}
	accessToken, _ = c.GetAccessToken()
	if accessToken != "1234" {
		t.Errorf("Unexpected Access Token. Expecting 1234, Got %s", accessToken)
	}

	c = Config{Environment: Production, SandboxAccessToken: "1234", ProductionAccessToken: "4567"}
	accessToken, _ = c.GetAccessToken()
	if accessToken != "4567" {
		t.Errorf("Unexpected Access Token. Expecting 4567, Got %s", accessToken)
	}
}
