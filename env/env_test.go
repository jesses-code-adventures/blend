package env

import (
	"os"
	"testing"
)

func Test_LoadEnvVars(t *testing.T) {
	test := true
	LoadEnvVars(test)
	ensure_vars := []string{"TEST_DATA_DIR"}
	for _, v := range ensure_vars {
		if os.Getenv(v) == "" {
			t.Errorf("Expected %s to be set", v)
		}
	}
}

func Test_TestPrecedenceVar(t *testing.T) {
	test := true
	LoadEnvVars(test)
	if os.Getenv("TEST_PRECEDENCE_VAR") != ".env.test" {
		t.Errorf("TEST_PRECEDENCE_VAR should be .env.test in when test is true")
	}
}
