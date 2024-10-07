package ingest

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jesses-code-adventures/blend/env"
)

func Test_UnixFilepathIngestor(t *testing.T) {
	env.LoadEnvVarsWithTestVars()
	testDataDir := os.Getenv("TEST_DATA_DIR")
	dir := fmt.Sprintf("../%s/sf_bash_in", testDataDir)
	i := UnixFilepathIngestorFromRoot(dir)
	i.Ingest()
	if len(i.Locations()) == 0 {
		t.Errorf("No test files in %s", dir)
	}
	for f, _ := range i.Locations() {
		if !strings.Contains(f, "hello_world") {
			t.Errorf("expected string containing, got %s", f)
		}
	}
}
