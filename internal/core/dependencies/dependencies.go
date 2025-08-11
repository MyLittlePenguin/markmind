package dependencies

import (
	"markmind/internal/data/repos"
	"markmind/internal/domain/usecases"
)

var fileRepo = repos.NewFileRepo()
var FileUseCase = usecases.NewFileUseCase(fileRepo)

var explorerRepo = repos.NewExplorerRepo()
var ExplorerUseCase = usecases.NewExplorerUseCase(explorerRepo)

var GraphUseCase = usecases.NewGraphUseCase(explorerRepo, fileRepo)
