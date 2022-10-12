package models

type Input struct {
	RootName string
	Post     RPost
	Posts    []RPost
	Tag      RTag
	Tags     []RTag
	Content  string
}

type RTag struct {
	Name  string
	Count int
	Path  string
	Posts []RPost
}

type RPost struct {
	Title   string
	URL     string
	Date    string
	Tags    []string
	Content string
}
