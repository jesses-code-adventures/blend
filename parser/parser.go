package parser

import (
	"bufio"
	"bytes"
	"io"
)

type FileChange struct {
	relativePath string
	lines        [][]byte
}

func (f FileChange) FileContents() bytes.Buffer {
	var resp bytes.Buffer
	for i, line := range f.lines {
		resp.Write(line)
		if i != len(f.lines)-1 {
			resp.WriteByte('\n')
		}
	}
	return resp
}

type Parser struct {
	changes []FileChange
	parsing bool
}

func NewParser() Parser {
	return Parser{
		changes: make([]FileChange, 0),
		parsing: false,
	}
}

func (p *Parser) parseLines(scanner *bufio.Scanner) {
	relativePath := ""
	lines := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		if p.parsing && !startsBackticks(line) {
			lines = append(lines, line)
			continue
		}
		if p.parsing && startsBackticks(line) {
			p.changes = append(p.changes, FileChange{
				relativePath: relativePath,
				lines:        lines,
			})
			relativePath = ""
			p.parsing = false
			lines = make([][]byte, 0)
			continue
		}
		if !p.parsing && startsBackticks(line) {
			p.parsing = true
			relativePath = string(line[3:])
			continue
		}
	}

}

func (p *Parser) Parse(reader io.ReadCloser) {
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	p.parseLines(scanner)
}

func (p Parser) ParsedFilesAsString() string {
	var buffer bytes.Buffer
	for _, change := range p.changes {
		buffer.WriteString(change.relativePath)
		buffer.WriteByte('\n')
		contents := change.FileContents()
		buffer.Write(contents.Bytes())
	}
	return buffer.String()
}

func startsBackticks(b []byte) bool {
	if len(b) < 3 {
		return false
	}
	return b[0] == '`' && b[1] == '`' && b[2] == '`'
}
