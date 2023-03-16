package app

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/qunv/minitino/app/extractor"
	"github.com/qunv/minitino/app/helpers"
	"github.com/qunv/minitino/app/models"
	"github.com/russross/blackfriday"
	"os"
	"strings"
	"text/template"
)

var tagMap = make(map[string]*models.RenderTag)
var rPosts []models.RenderPost
var posts []models.ExtractedPost

type app struct {
	ctx           context.Context
	fs            embed.FS
	postExtractor extractor.Extractor[[]models.ExtractedPost]
	config        models.Config
}

func getRenderTagsByKeys(tags []string) []models.RenderTag {
	var resp []models.RenderTag
	for _, t := range tags {
		if rt, ok := tagMap[t]; ok {
			resp = append(resp, *rt)
		}
	}
	return resp
}

func (a app) Run() {
	a.makeDir()

	posts = a.postExtractor.Extract()
	for _, p := range posts {
		url := a.buildPostUrl(p.CreatedAt, p.Title)
		a.handleTag(p.Tags, url, p.CreatedAt, p.Title)
		rPosts = append(rPosts, models.RenderPost{
			BasePost: models.BasePost{
				Title:     p.Title,
				CreatedAt: helpers.ConvertDate(p.CreatedAt),
			},
			URL:  url,
			Tags: getRenderTagsByKeys(p.Tags),
		})
	}

	a.renderAssets(models.SysAssetsDir, models.AssetsDir)
	a.renderIndexPage()
	a.renderPostPages()
	a.renderTagsPage()
	a.renderTagDetailPage()
	a.renderAbout()
	a.renderRSS()
	a.renderPoem()
}

func (a app) handleTag(tags []string, path, date, title string) {
	for _, tag := range tags {
		if curTag, ok := tagMap[tag]; ok {
			curTag.Name = tag
			curTag.Count++
			curTag.Posts = append(curTag.Posts, models.RenderPost{
				URL: path,
				BasePost: models.BasePost{
					CreatedAt: date,
					Title:     title,
				},
				RawTags: tags,
			})
		} else {
			tagPath := fmt.Sprintf("tags/%s", tag)
			_ = os.MkdirAll(tagPath, 0755)
			tagMap[tag] = &models.RenderTag{
				Path:     "/" + tagPath,
				Name:     tag,
				Count:    1,
				ColorHEX: helpers.RandomTagColors(len(tagMap)),
				Posts: []models.RenderPost{
					{
						BasePost: models.BasePost{
							Title:     title,
							CreatedAt: date,
						},
						URL:     path,
						RawTags: tags,
					},
				},
			}
		}
	}
}

func (a app) makeDir() {
	_ = os.RemoveAll(models.PostsDir)

	_ = os.MkdirAll(models.SysPostsDir, 0755)
	_ = os.MkdirAll(models.SysAboutDir, 0755)
	_ = os.MkdirAll(models.SysPoemDir, 0755)

	_ = os.MkdirAll(models.PostsDir, 0755)
	_ = os.MkdirAll(models.AssetsDir, 0755)
	_ = os.MkdirAll(models.AboutDir, 0755)
	_ = os.MkdirAll(models.ImagesDir, 0755)
	_ = os.MkdirAll(models.TagsDir, 0755)
	_ = os.MkdirAll(models.PoemDir, 0755)
}

func (a app) renderAssets(sysDir string, makeDir string) {
	dirs, err := a.fs.ReadDir(sysDir)
	helpers.PanicIfError(err)
	for _, dir := range dirs {
		sys := sysDir + "/" + dir.Name()
		md := makeDir + "/" + dir.Name()
		if dir.IsDir() {
			_ = os.MkdirAll(md, 0755)
			a.renderAssets(sys, md)
			continue
		}
		file, err := a.fs.ReadFile(sys)
		helpers.PanicIfError(err)
		err = helpers.WriteFile(md, bytes.NewBuffer(file))
	}
}

func (a app) renderIndexPage() {
	t, err := template.ParseFS(
		a.fs,
		"_templates/root.gohtml",
		"_templates/index/indexBody.gohtml",
		"_templates/index/indexSubHeader.gohtml",
	)
	helpers.PanicIfError(err)
	b := &bytes.Buffer{}

	data := models.Input{
		Config: a.config,
		Posts:  rPosts,
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
		t, err := template.ParseFS(
			a.fs,
			"_templates/root.gohtml",
			"_templates/post/postBody.gohtml",
			"_templates/post/postSubHeader.gohtml",
		)
		helpers.PanicIfError(err)

		b := &bytes.Buffer{}

		file, err := helpers.ReadFile(p.FilePath)
		helpers.PanicIfError(err)

		pars := blackfriday.MarkdownCommon(file.Bytes())

		input := models.Input{
			Config: a.config,
			Post: models.RenderPost{
				BasePost: models.BasePost{
					Title:     p.Title,
					CreatedAt: p.CreatedAt,
				},
				Tags:    getRenderTagsByKeys(p.Tags),
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

func (a app) renderAbout() {
	t, err := template.ParseFS(
		a.fs,
		"_templates/root.gohtml",
		"_templates/about/aboutBody.gohtml",
		"_templates/about/aboutSubHeader.gohtml",
	)
	helpers.PanicIfError(err)

	b := &bytes.Buffer{}

	file, err := helpers.ReadFile("_about/about.md")
	if err != nil {
		return
	}

	pars := blackfriday.MarkdownCommon(file.Bytes())

	input := models.Input{
		Config:  a.config,
		Content: string(pars),
	}

	err = t.Execute(b, input)
	if err != nil {
		panic(err)
	}
	err = helpers.WriteFile("about/index.html", b)
	helpers.PanicIfError(err)
}

func (a app) renderPoem() {
	t, err := template.ParseFS(
		a.fs,
		"_templates/root.gohtml",
		"_templates/poem/poemBody.gohtml",
		"_templates/poem/poemSubHeader.gohtml",
	)
	helpers.PanicIfError(err)

	b := &bytes.Buffer{}

	file, err := helpers.ReadFile("_poem/poem.md")
	if err != nil {
		return
	}

	pars := blackfriday.MarkdownCommon(file.Bytes())

	input := models.Input{
		Config:  a.config,
		Content: string(pars),
	}

	err = t.Execute(b, input)
	if err != nil {
		panic(err)
	}
	err = helpers.WriteFile("poem/index.html", b)
	helpers.PanicIfError(err)
}

func (a app) renderTagsPage() {
	t, err := template.ParseFS(
		a.fs,
		"_templates/root.gohtml",
		"_templates/tag/tagListBody.gohtml",
		"_templates/tag/tagListSubHeader.gohtml",
	)
	helpers.PanicIfError(err)
	b := &bytes.Buffer{}
	var tags []models.RenderTag
	for _, tag := range tagMap {
		tags = append(tags, *tag)
	}
	data := models.Input{
		Config: a.config,
		Tags:   tags,
	}
	err = t.Execute(b, data)
	if err != nil {
		panic(err)
	}
	err = helpers.WriteFile("tags/index.html", b)
	helpers.PanicIfError(err)
}

func (a app) renderTagDetailPage() {
	for tagName, tag := range tagMap {
		t, err := template.ParseFS(
			a.fs,
			"_templates/root.gohtml",
			"_templates/tag/tagBody.gohtml",
			"_templates/tag/tagSubHeader.gohtml",
		)
		helpers.PanicIfError(err)

		b := &bytes.Buffer{}

		var ps []models.RenderPost

		for _, post := range tag.Posts {
			post.Tags = getRenderTagsByKeys(post.RawTags)
			ps = append(ps, post)
		}

		tag.Posts = ps

		input := models.Input{
			Config: a.config,
			Tag:    *tag,
		}

		err = t.Execute(b, input)
		if err != nil {
			panic(err)
		}
		err = helpers.WriteFile("tags/"+tagName+"/index.html", b)
		helpers.PanicIfError(err)
	}
}

func (a app) renderRSS() {
	var b bytes.Buffer
	b.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	b.WriteString("<rss version=\"2.0\" xmlns:atom=\"http://www.w3.org/2005/Atom\">\n")
	b.WriteString("<channel>\n")
	b.WriteString("<title>" + a.config.RootName + "</title>\n")
	b.WriteString("<description>For the Future</description>\n")
	b.WriteString("<link>https://qunv.github.io/</link>\n")
	for _, post := range posts {
		dateFolder := strings.ReplaceAll(post.CreatedAt, "-", "/")
		path := "posts/" + dateFolder + "/" + strings.ReplaceAll(post.Title, " ", "-")
		b.WriteString("<item>\n")
		b.WriteString("<title>" + post.Title + "</title>\n")
		b.WriteString("<link>https://qunv.github.io/" + path + "</link>\n")
		b.WriteString("</item>\n")
	}
	b.WriteString("</channel>\n")
	b.WriteString("</rss>\n")
	err := helpers.WriteFile("rss.xml", &b)
	helpers.PanicIfError(err)
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
