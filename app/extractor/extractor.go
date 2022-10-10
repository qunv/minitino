package extractor

import (
	"bytes"
	"github.com/qunv/minitino/app/helpers"
	"github.com/qunv/minitino/app/models"
	"log"
	"regexp"
	"strings"
)

type Extractor[O any] interface {
	Extract() O
}

func NewPostExtractor(source string) Extractor[[]models.Post] {
	return &postExtractor{
		dirPath: source,
	}
}

type postExtractor struct {
	dirPath string
}

func (p postExtractor) Extract() []models.Post {
	dirs, err := helpers.ReadDir(p.dirPath)
	helpers.PanicIfError(err)
	var posts []models.Post
	for _, dir := range dirs {
		fileName := dir.Name()
		file, err := helpers.ReadFile(p.dirPath + "/" + fileName)
		helpers.PanicIfError(err)
		createdAt := fileName[0:10]
		title := p.extractTitle(file)
		tags := p.extractTags(file)

		posts = append(posts, models.Post{
			Content:   file.Bytes(),
			Title:     title,
			Tags:      tags,
			CreatedAt: createdAt,
		})
	}
	return posts
}

func (p postExtractor) extractTitle(file *bytes.Buffer) string {
	fileBytes := file.Bytes()

	title := strings.Split(string(fileBytes), "\n")[0]
	match, err := regexp.MatchString("^\\[title\\]:<>\\(.*\\)$", title)
	helpers.PanicIfError(err)
	if !match {
		log.Println("Missing title!")
		panic("Missing title")
	}

	title = strings.ReplaceAll(title, "[title]:<>(", "")
	title = strings.ReplaceAll(title, ")", "")
	return title
}

func (p postExtractor) extractTags(file *bytes.Buffer) []string {
	fileBytes := file.Bytes()

	tags := strings.Split(string(fileBytes), "\n\n")[1]
	match, err := regexp.MatchString("^\\[tags\\]:<>\\(.*\\)$", tags)
	helpers.PanicIfError(err)

	if !match {
		log.Println("Missing tags format!")
		panic("Missing tags")
	}

	tags = strings.ReplaceAll(tags, "[tags]:<>(", "")
	tags = strings.ReplaceAll(tags, ")", "")
	return strings.Split(tags, ",")
}
