package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil // found go.mod, this is the root
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", os.ErrNotExist
}

func globAndOverrideVars(pattern string) (matches []string) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	return matches
}

func loadFromGlob(matches []string) {
	for _, match := range matches {
		err := godotenv.Overload(match)
		if err != nil {
			log.Fatalf("Error loading %s file: %v", match, err)
		}
		fmt.Sprintf("loaded %s", match)
	}
	return
}

func LoadEnvVars() {
	workingDir, err := findProjectRoot()
	if err != nil {
		panic(err)
	}
	matches := globAndOverrideVars(fmt.Sprintf("%s/*.env.public", workingDir))
	loadFromGlob(matches)
	matches = globAndOverrideVars(fmt.Sprintf("%s/*.env.mine", workingDir))
	loadFromGlob(matches)
}

func LoadEnvVarsWithTestVars() {
	repoRoot, err := findProjectRoot()
	if err != nil {
		panic(err)
	}
	matches := globAndOverrideVars(fmt.Sprintf("%s/*.env.public", repoRoot))
	loadFromGlob(matches)
	matches = globAndOverrideVars(fmt.Sprintf("%s/*.env.mine", repoRoot))
	loadFromGlob(matches)
	testPath := fmt.Sprintf("%s/*.env.test", repoRoot)
	matches = globAndOverrideVars(testPath)
	loadFromGlob(matches)
}
