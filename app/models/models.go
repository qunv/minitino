package models

import "bytes"

type Post struct {
	FilePath  string
	Title     string
	CreatedAt string
	Tags      []string
	Content   []byte
}

type TemplateInfo struct {
	Path    string
	Content *bytes.Buffer
}
