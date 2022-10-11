package models

type Index struct {
	RootName string
	Posts    []RPost
}

type PostInput struct {
	RootName string
	Post     RPost
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
