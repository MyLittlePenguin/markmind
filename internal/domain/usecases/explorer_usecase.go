package usecases

import (
	"markmind/internal/core/utils"
	"markmind/internal/data/entities"
	"markmind/internal/data/repos"
	"strings"
)

type ExplorerUseCase struct {
	repo *repos.ExplorerRepo
}

func NewExplorerUseCase(repo *repos.ExplorerRepo) *ExplorerUseCase {
	return &ExplorerUseCase{
		repo: repo,
	}
}

func (self *ExplorerUseCase) GetEntries() ([]entities.MarkdownFileMeta, error) {
	return self.repo.GetEntries()
}

func (self *ExplorerUseCase) GetEntriesOfDirectory(directory string) ([]entities.MarkdownFileMeta, error) {
	return self.repo.GetEntriesOfDirectory(directory)
}


func (self *ExplorerUseCase) MakeDir(path string, name string) error {
	path = strings.TrimSuffix(path, "/")
  newPath, err := utils.SanitizePath(path + "/" + name)
  if err != nil {
    return err
  }
	return self.repo.MakeDir(newPath)
}

func (self *ExplorerUseCase) Delete(path string) error {
  newPath, err := utils.SanitizePath(path)
  if err != nil {
    return err
  }
  return self.repo.Delete(newPath)
}
