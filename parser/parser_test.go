package parser

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test_ParserHappy(t *testing.T) {
	tokens := "```test_data/sf_bash_in/hello_world.bash\n#!/bin/bash\n\nehco 'hello world!'"
	numLines := 3
	reader := io.NopCloser(strings.NewReader(tokens))
	parser := NewParser()
	err := parser.Parse(reader)
	if err != nil {
		t.Errorf("parsing failed")
	}
	for _, change := range parser.changes {
		fmt.Println(change.relativePath)
		if change.relativePath == "" {
			t.Errorf("no relative path")
		}
		if len(change.lines) == 0 {
			t.Errorf("failed to parse lines")
		}
		if len(change.lines) < numLines {
			t.Errorf("parsed too few lines")
		}
		if len(change.lines) > numLines {
			t.Errorf("parsed too many lines")
		}
	}
}
