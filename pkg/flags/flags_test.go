package flags

import "testing"

func TestSandboxEnv(t *testing.T) {
	ef := &EnvironmentFlag{
		env: "production",
	}

	if ef.String() != "production" {
		t.Errorf("Expected production, got %s", ef.String())
	}
}

func TestNoProvidedEnv(t *testing.T) {
	ef := &EnvironmentFlag{}

	if ef.String() != "sandbox" {
		t.Errorf("Expected sandbox, got %s", ef.String())
	}
}

func TestInvalidEnv(t *testing.T) {
	ef := &EnvironmentFlag{}

	if ef.Set("foobar") == nil {
		t.Errorf("Expected error when setting invalid env")
	}
}
