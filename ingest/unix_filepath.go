package ingest

import (
	"os"
	"path/filepath"
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
	filepaths, err := filepath.Glob(filepath.Join(ufi.root, "*"))
	if err != nil {
		return
	}
	for _, path := range filepaths {
		ufi.paths[path] = struct{}{}
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		ufi.contents[path] = string(content)
	}
}

func (ufi *UnixFilepathIngestor) Locations() map[string]struct{} {
	return ufi.paths
}

func (ufi *UnixFilepathIngestor) Contents() map[string]string {
	return ufi.contents
}
