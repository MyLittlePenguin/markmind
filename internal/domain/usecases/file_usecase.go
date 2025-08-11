package usecases

import (
	"errors"
	"markmind/internal/core/moner"
	"markmind/internal/core/utils"
	"markmind/internal/data/repos"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type FileUseCase struct {
	repo *repos.FileRepo
}

func NewFileUseCase(repo *repos.FileRepo) *FileUseCase {
	return &FileUseCase{
		repo: repo,
	}
}

func (self *FileUseCase) CreateFile(path string, name string) error {
	err := utils.PathIsOk(path)
	if err != nil {
		return err
	}
	err = utils.PathIsOk(name)
	if err != nil {
		return err
	}

	if strings.Contains(name, "/") {
		return errors.New("No '/' allowed inside file name")
	}
	_, err = self.repo.CreateFile(path, name)
	return err
}

func (self *FileUseCase) GetFileContent(file string) (string, error) {
	return moner.Fmap(
		mdToHTML,
		self.repo.GetFileContent,
	)(file)
}

func (self *FileUseCase) GetRawFileContent(file string) (string, error) {
	return moner.Fmap(
		toString,
		self.repo.GetFileContent,
	)(file)
}

func (self *FileUseCase) UpdateFileContent(path string, content *string) error {
	sanitizedPath, err := utils.SanitizePath(path)
	if err != nil {
		return err
	}
	return self.repo.UpdateFileContent(sanitizedPath, content)
}

func (self *FileUseCase) GetParentDir(path string) (string, error) {
	return moner.Bind(
		utils.SanitizePath,
		self.repo.GetParentDir,
	)(path)
}

func (self *FileUseCase) IsDir(path string) (bool, error) {
	return moner.Bind(
		utils.SanitizePath,
		self.repo.IsDir,
	)(path)
}

func mdToHTML(md []byte) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	mdParser := parser.NewWithExtensions(extensions)
	docAst := mdParser.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(docAst, renderer))
}

func toString(it []byte) string {
	return string(it)
}
