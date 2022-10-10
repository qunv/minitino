package app

import (
	"context"
	"github.com/qunv/minitino/app/extractor"
	helpers2 "github.com/qunv/minitino/app/helpers"
	"github.com/qunv/minitino/app/models"
	"html/template"
	"os"
)

type app struct {
	ctx           context.Context
	templates     map[string]models.TemplateInfo
	postExtractor extractor.Extractor[[]models.Post]
}

func (a app) Run() {
	a.setup()
	a.renderAssets()
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
	dirs, err := helpers2.ReadDir(models.SysAssetsDir)
	helpers2.PanicIfError(err)
	for _, dir := range dirs {
		sysFileName := models.SysAssetsDir + "/" + dir.Name()
		file, err := helpers2.ReadFile(sysFileName)
		helpers2.PanicIfError(err)
		err = helpers2.WriteFile(models.AssetsDir+"/"+dir.Name(), file)
	}
}

func (a app) renderIndex() {
	if templateInfo, ok := a.templates["index.html"]; ok {
		var
		t := template.Must(template.New("html-tmpl").ParseFiles(templateInfo.Path))
		err := t.Execute(os.Stdout, )
		if err != nil {
			panic(err)
		}
	}
}
