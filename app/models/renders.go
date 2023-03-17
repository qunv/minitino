package models

import "bytes"

type BasePost struct {
	Title       string
	CreatedAt   string
	Description string
}

type ExtractedPost struct {
	BasePost
	FilePath string
	Tags     []string
	Raw      []byte
}

type RenderPost struct {
	BasePost
	URL     string
	RawTags []string
	Tags    []RenderTag
	Content string
}

type RenderTag struct {
	Name     string
	Count    int
	Path     string
	ColorHEX string
	Posts    []RenderPost
}

type TemplateInfo struct {
	Path    string
	Content *bytes.Buffer
}

type Input struct {
	Config  Config
	Post    RenderPost
	Posts   []RenderPost
	Tag     RenderTag
	Tags    []RenderTag
	Content string
}
