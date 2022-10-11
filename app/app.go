package app

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qunv/minitino/app/extractor"
	"github.com/qunv/minitino/app/helpers"
	"github.com/qunv/minitino/app/models"
	"github.com/russross/blackfriday"
	"os"
	"strings"
	"text/template"
)

var tagMap = make(map[string]*models.RTag)
var rPosts []models.RPost
var posts []models.Post

type app struct {
	ctx           context.Context
	templates     map[string]models.TemplateInfo
	postExtractor extractor.Extractor[[]models.Post]
}

func (a app) Run() {
	a.setup()

	posts = a.postExtractor.Extract()
	for _, p := range posts {
		url := a.buildPostUrl(p.CreatedAt, p.Title)
		a.handleTag(p.Tags, url, p.CreatedAt, p.Title)
		rPosts = append(rPosts, models.RPost{
			URL:   url,
			Title: p.Title,
			Tags:  p.Tags,
			Date:  helpers.ConvertDate(p.CreatedAt),
		})
	}

	a.renderAssets()
	a.renderIndexPage()
	a.renderPostPages()
}

func (a app) handleTag(tags []string, path, date, title string) {
	for _, tag := range tags {
		if val, ok := tagMap[tag]; ok {
			val.Name = tag
			val.Count++
			val.Posts = append(val.Posts, models.RPost{
				URL:   path,
				Date:  date,
				Title: title,
				Tags:  tags,
			})
		} else {
			tagPath := fmt.Sprintf("tags/%s", tag)
			_ = os.MkdirAll(tagPath, 0755)
			tagMap[tag] = &models.RTag{
				Path:  tagPath,
				Name:  tag,
				Count: 1,
				Posts: []models.RPost{
					{
						URL:   path,
						Date:  date,
						Title: title,
						Tags:  tags,
					},
				},
			}
		}
	}
}

func (a app) setup() {
	_ = os.RemoveAll(models.PostsDir)

	_ = os.MkdirAll(models.SysSectionsDir, 0755)
	_ = os.MkdirAll(models.SysPostsDir, 0755)
	_ = os.MkdirAll(models.SysAboutDir, 0755)

	_ = os.MkdirAll(models.PostsDir, 0755)
	_ = os.MkdirAll(models.AssetsDir, 0755)
	_ = os.MkdirAll(models.AboutDir, 0755)
	_ = os.MkdirAll(models.ImagesDir, 0755)
	_ = os.MkdirAll(models.TagsDir, 0755)
}

func (a app) renderAssets() {
	dirs, err := helpers.ReadDir(models.SysAssetsDir)
	helpers.PanicIfError(err)
	for _, dir := range dirs {
		sysFileName := models.SysAssetsDir + "/" + dir.Name()
		file, err := helpers.ReadFile(sysFileName)
		helpers.PanicIfError(err)
		err = helpers.WriteFile(models.AssetsDir+"/"+dir.Name(), file)
	}
}

func (a app) renderIndexPage() {
	indexTpl := a.templates["index.html"]
	postListBodyTpl := a.templates["postListBody.html"]
	indexSubHeaderTpl := a.templates["indexSubHeader.html"]
	t, err := template.ParseFiles(indexTpl.Path, postListBodyTpl.Path, indexSubHeaderTpl.Path)
	helpers.PanicIfError(err)
	b := &bytes.Buffer{}

	data := models.Index{
		RootName: "JUANTINO NG",
		Posts:    rPosts,
	}
	err = t.Execute(b, data)
	if err != nil {
		panic(err)
	}
	err = helpers.WriteFile("index.html", b)
	helpers.PanicIfError(err)
}

func (a app) renderPostPages() {
	for _, p := range posts {
		indexTpl := a.templates["index.html"]
		postSubHeaderTpl := a.templates["postSubHeader.html"]
		postBodyTpl := a.templates["postBody.html"]
		t, err := template.ParseFiles(indexTpl.Path, postBodyTpl.Path, postSubHeaderTpl.Path)
		helpers.PanicIfError(err)

		b := &bytes.Buffer{}

		file, err := helpers.ReadFile(p.FilePath)
		helpers.PanicIfError(err)

		pars := blackfriday.MarkdownCommon(file.Bytes())

		input := models.PostInput{
			RootName: "JUANTINO NG",
			Post: models.RPost{
				Title:   p.Title,
				Tags:    p.Tags,
				Date:    helpers.ConvertDate(p.CreatedAt),
				Content: string(pars),
			},
		}

		err = t.Execute(b, input)
		if err != nil {
			panic(err)
		}
		a.createPostIndexFile(p.CreatedAt, p.Title, b)
	}
}

func (a app) createPostIndexFile(date, title string, b *bytes.Buffer) {
	dateFolder := strings.ReplaceAll(date, "-", "/")
	dir := "posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
	_ = os.MkdirAll(dir, 0755)
	err := helpers.WriteFile(dir+"/index.html", b)
	helpers.PanicIfError(err)
}

func (a app) buildPostUrl(date string, title string) string {
	dateFolder := strings.ReplaceAll(date, "-", "/")
	return "/posts/" + dateFolder + "/" + strings.ReplaceAll(title, " ", "-")
}
