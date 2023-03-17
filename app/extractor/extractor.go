package extractor

import (
	"bytes"
	"github.com/qunv/minitino/app/helpers"
	"github.com/qunv/minitino/app/models"
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
		createdAt := p.extractDate(file)
		title := p.extractTitle(file)
		tags := p.extractTags(file)
		description := p.extractDescription(file)

		posts = append(posts, models.ExtractedPost{
			BasePost: models.BasePost{
				Title:       title,
				CreatedAt:   helpers.ConvertDate(createdAt),
				Description: description,
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

	r := regexp.MustCompile(`\[title\]: <> \(.*?\)`)
	title := r.FindString(string(fileBytes))
	title = strings.ReplaceAll(title, "[title]: <> (", "")
	title = strings.ReplaceAll(title, ")", "")
	return title
}

func (p postExtractor) extractDescription(file *bytes.Buffer) string {
	fileBytes := file.Bytes()

	r := regexp.MustCompile(`\[description\]: <> \(.*?\)`)
	title := r.FindString(string(fileBytes))
	title = strings.ReplaceAll(title, "[description]: <> (", "")
	title = strings.ReplaceAll(title, ")", "")
	return title
}

func (p postExtractor) extractDate(file *bytes.Buffer) string {
	fileBytes := file.Bytes()

	r := regexp.MustCompile(`\[date\]: <> \(.*?\)`)
	title := r.FindString(string(fileBytes))
	title = strings.ReplaceAll(title, "[date]: <> (", "")
	title = strings.ReplaceAll(title, ")", "")
	return title
}

func (p postExtractor) extractTags(file *bytes.Buffer) []string {
	fileBytes := file.Bytes()

	r := regexp.MustCompile(`\[tags\]: <> \(.*?\)`)
	tags := r.FindString(string(fileBytes))

	tags = strings.ReplaceAll(tags, "[tags]: <> (", "")
	tags = strings.ReplaceAll(tags, ")", "")

	return strings.Split(tags, ",")
}
