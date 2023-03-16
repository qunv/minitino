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

func NewPostExtractor(source string) Extractor[[]models.ExtractedPost] {
	return &postExtractor{
		dirPath: source,
	}
}

type postExtractor struct {
	dirPath string
}

func (p postExtractor) Extract() []models.ExtractedPost {
	dirs, err := helpers.ReadDir(p.dirPath)
	helpers.PanicIfError(err)
	var posts []models.ExtractedPost
	for i := len(dirs) - 1; i >= 0; i-- {
		dir := dirs[i]
		fileName := dir.Name()
		filePath := p.dirPath + "/" + fileName
		file, err := helpers.ReadFile(filePath)
		helpers.PanicIfError(err)
		createdAt := fileName[0:10]
		title := p.extractTitle(file)
		tags := p.extractTags(file)

		posts = append(posts, models.ExtractedPost{
			BasePost: models.BasePost{
				Title:     title,
				CreatedAt: createdAt,
			},
			FilePath: filePath,
			Raw:      file.Bytes(),
			Tags:     tags,
		})
	}
	return posts
}

func (p postExtractor) extractTitle(file *bytes.Buffer) string {
	fileBytes := file.Bytes()

	title := strings.Split(string(fileBytes), "\n")[0]
	match, err := regexp.MatchString("^\\[comment\\]: <> \\(.*\\)$", title)
	helpers.PanicIfError(err)
	if !match {
		log.Println("Missing title!")
		panic("Missing title")
	}

	title = strings.ReplaceAll(title, "[comment]: <> (", "")
	title = strings.ReplaceAll(title, ")", "")
	return title
}

func (p postExtractor) extractTags(file *bytes.Buffer) []string {
	fileBytes := file.Bytes()

	tags := strings.Split(string(fileBytes), "\n\n")[1]
	match, err := regexp.MatchString("^\\[comment\\]: <> \\(.*\\)$", tags)
	helpers.PanicIfError(err)

	if !match {
		log.Println("Missing tags format!")
		panic("Missing tags")
	}

	tags = strings.ReplaceAll(tags, "[comment]: <> (", "")
	tags = strings.ReplaceAll(tags, ")", "")
	return strings.Split(tags, ",")
}
