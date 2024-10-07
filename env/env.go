package env

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func globAndOverrideVars(pattern string) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, match := range matches {
		err := godotenv.Overload(match)
		if err != nil {
			log.Fatalf("Error loading %s file: %v", match, err)
		}
		fmt.Sprintf("loaded %s", match)
	}
	return
}

func LoadEnvVars(test bool) {
	globAndOverrideVars("../*.env.public")
	globAndOverrideVars("../*.env.mine")
	if test {
		globAndOverrideVars("../*.env.test")
	}
}
