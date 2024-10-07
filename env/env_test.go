package env

import (
	"os"
	"testing"
)

func Test_LoadEnvVars(t *testing.T) {
	LoadEnvVarsWithTestVars()
	ensure_vars := []string{"TEST_PRECEDENCE_VAR"}
	for _, v := range ensure_vars {
		if os.Getenv(v) == "" {
			t.Errorf("Expected %s to be set", v)
		}
	}
}

func Test_LoadEnvVarsWithTestVars(t *testing.T) {
	LoadEnvVarsWithTestVars()
	ensure_vars := []string{"TEST_DATA_DIR"}
	for _, v := range ensure_vars {
		if os.Getenv(v) == "" {
			t.Errorf("Expected %s to be set", v)
		}
	}
}

func Test_TestPrecedenceVar(t *testing.T) {
	LoadEnvVarsWithTestVars()
	if os.Getenv("TEST_PRECEDENCE_VAR") != ".env.test" {
		t.Errorf("TEST_PRECEDENCE_VAR should be .env.test in when test is true, got %s", os.Getenv("TEST_PRECEDENCE_VAR"))
	}
}

func Test_TestPrecedenceVarNonTest(t *testing.T) {
	LoadEnvVars()
	if os.Getenv("TEST_PRECEDENCE_VAR") != ".env.mine" {
		t.Errorf("TEST_PRECEDENCE_VAR should be .env.mine in when test is false, got %s", os.Getenv("TEST_PRECEDENCE_VAR"))
	}
}

func Test_OpenaiApiKey(t *testing.T) {
	LoadEnvVarsWithTestVars()
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Errorf("Expected to find an OPENAI_API_KEY")
	}
}
