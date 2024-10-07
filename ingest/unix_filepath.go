package ingest

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type UnixFilepathIngestor struct {
	root     string
	paths    map[string]struct{}
	contents map[string]string
}

func NewUnixFilepathIngestor() *UnixFilepathIngestor {
	return &UnixFilepathIngestor{
		root:     ".",
		paths:    make(map[string]struct{}),
		contents: make(map[string]string),
	}
}

func UnixFilepathIngestorFromRoot(root string) *UnixFilepathIngestor {
	return &UnixFilepathIngestor{
		root:     root,
		paths:    make(map[string]struct{}),
		contents: make(map[string]string),
	}
}

func (ufi *UnixFilepathIngestor) Ingest() {
	// TODO: make this use os.Getenv("FILETYPES")
	ufi.paths = make(map[string]struct{})
	ufi.contents = make(map[string]string)
	err := filepath.WalkDir(ufi.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".bash" {
			ufi.paths[path] = struct{}{}
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			ufi.contents[path] = string(content)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through directories:", err)
	}
}

func (ufi *UnixFilepathIngestor) Locations() map[string]struct{} {
	return ufi.paths
}

func (ufi *UnixFilepathIngestor) Contents() map[string]string {
	return ufi.contents
}

func (ufi *UnixFilepathIngestor) ContentsString() string {
	var b strings.Builder
	for path, contents := range ufi.Contents() {
		b.WriteString(fmt.Sprintf("```%s\n%s\n```\n", path, contents))
	}
	return b.String()
}
