package boot

import (
	"bytes"
	"context"
	"github.com/qunv/minitino/helpers"
	"os"
)

type Boot interface {
	Run()
}

func New(ctx context.Context) Boot {
	return &app{
		ctx: ctx,
	}
}

//const (
//	SysAssetsDir   string = "_assets"
//	SysSectionsDir        = "_sections"
//	SysPostsDir           = "_posts"
//	SysAboutDir           = "_about"
//)
//
//const (
//	AssetsDir   string = "assets"
//	SectionsDir        = "sections"
//	PostsDir           = "posts"
//	AboutDir           = "about"
//	ImagesDir          = "images"
//	TagsDir            = "tags"
//)

const (
	SysAssetsDir    string = "_assets"
	SysSectionsDir         = "example/_sections"
	SysPostsDir            = "example/_posts"
	SysAboutDir            = "example/_about"
	SysTemplatesDir        = "_templates"
)

const (
	AssetsDir   string = "example/assets"
	SectionsDir        = "example/sections"
	PostsDir           = "example/posts"
	AboutDir           = "example/about"
	ImagesDir          = "example/images"
	TagsDir            = "example/tags"
)

func initTemplates() map[string]*bytes.Buffer {
	dirs, err := helpers.ReadDir(SysTemplatesDir)
	helpers.Panic(err)
	resp := make(map[string]*bytes.Buffer)
	for _, dir := range dirs {
		fileName := dir.Name()
		filePath := SysTemplatesDir + "/" + fileName
		file, err := helpers.ReadFile(filePath)
		helpers.Panic(err)
		
	}
	return resp
}

type app struct {
	ctx       context.Context
	templates map[string]string
}

func (a app) Run() {
	a.setup()
	a.initAssets()
}

func (a app) setup() {
	_ = os.RemoveAll(PostsDir)

	_ = os.MkdirAll(SysSectionsDir, 0755)
	_ = os.MkdirAll(SysPostsDir, 0755)
	_ = os.MkdirAll(SysAboutDir, 0755)

	_ = os.MkdirAll(PostsDir, 0755)
	_ = os.MkdirAll(AssetsDir, 0755)
	_ = os.MkdirAll(AboutDir, 0755)
	_ = os.MkdirAll(ImagesDir, 0755)
	_ = os.MkdirAll(TagsDir, 0755)
}

func (a app) initAssets() {
	dirs, err := helpers.ReadDir(SysAssetsDir)
	helpers.Panic(err)
	for _, dir := range dirs {
		sysFileName := SysAssetsDir + "/" + dir.Name()
		file, err := helpers.ReadFile(sysFileName)
		helpers.Panic(err)
		err = helpers.WriteFile(AssetsDir+"/"+dir.Name(), file)
	}
}

func (a app) parseIndex() {
}
