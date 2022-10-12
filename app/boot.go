package app

import (
	"context"
	"embed"
	"github.com/qunv/minitino/app/config"
	"github.com/qunv/minitino/app/extractor"
	"github.com/qunv/minitino/app/models"
)

type Boot interface {
	Run()
}

func New(ctx context.Context, fs embed.FS, config config.Config) Boot {
	postE := extractor.NewPostExtractor(models.SysPostsDir)
	return &app{
		ctx:           ctx,
		fs:            fs,
		postExtractor: postE,
		config:        config,
	}
}
