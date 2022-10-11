package app

import (
	"context"
	"github.com/qunv/minitino/app/extractor"
	helpers "github.com/qunv/minitino/app/helpers"
	"github.com/qunv/minitino/app/models"
)

type Boot interface {
	Run()
}

func New(ctx context.Context) Boot {
	templates := initTemplates()
	postE := extractor.NewPostExtractor(models.SysPostsDir)
	return &app{
		ctx:           ctx,
		templates:     templates,
		postExtractor: postE,
	}
}

func initTemplates() map[string]models.TemplateInfo {
	dirs, err := helpers.ReadDir(models.SysTemplatesDir)
	helpers.PanicIfError(err)
	resp := make(map[string]models.TemplateInfo)
	for _, dir := range dirs {
		fileName := dir.Name()
		filePath := models.SysTemplatesDir + "/" + fileName
		file, err := helpers.ReadFile(filePath)
		helpers.PanicIfError(err)
		resp[fileName] = models.TemplateInfo{
			Path:    models.SysTemplatesDir + "/" + fileName,
			Content: file,
		}
	}
	return resp
}
