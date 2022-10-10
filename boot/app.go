package boot

import (
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

const (
	SysAssetsDir   string = "_assets"
	SysSectionsDir        = "_sections"
	SysPostsDir           = "_posts"
	SysAboutDir           = "_about"
)

const (
	AssetsDir   string = "assets"
	SectionsDir        = "sections"
	PostsDir           = "posts"
	AboutDir           = "about"
	ImagesDir          = "images"
	TagsDir            = "tags"
)

type app struct {
	ctx context.Context
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
