package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

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
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	matches := globAndOverrideVars(fmt.Sprintf("%s/*.env.public", workingDir))
	loadFromGlob(matches)
	matches = globAndOverrideVars(fmt.Sprintf("%s/*.env.mine", workingDir))
	loadFromGlob(matches)
}

func LoadEnvVarsWithTestVars() {
	repoRoot, err := os.Getwd() // Change this to get the repository root
	if err != nil {
		panic(err)
	}
	matches := globAndOverrideVars(fmt.Sprintf("%s/*.env.public", repoRoot))
	loadFromGlob(matches)
	matches = globAndOverrideVars(fmt.Sprintf("%s/*.env.mine", repoRoot))
	loadFromGlob(matches)
	testPath := fmt.Sprintf("%s/*.env.test", repoRoot)
	// fmt.Print(testPath)
	matches = globAndOverrideVars(testPath)
	loadFromGlob(matches)
}
