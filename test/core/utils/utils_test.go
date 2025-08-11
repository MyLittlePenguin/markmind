package utils

import (
	"markmind/internal/core/utils"
	"testing"
)

func TestPathIsOk_noSurroundingSlashes(t *testing.T) {
  path := "a/b/c/d"
  err := utils.PathIsOk(path)
  if err != nil {
    t.Fatalf("For '%s' utils.PathIsOk returned the error: %s", path, err.Error())
  }
}

func TestPathIsOk_withSurroundingSlashes(t *testing.T) {
  path := "/a/b/c/d/"
  err := utils.PathIsOk(path)
  if err != nil {
    t.Fatalf("For '%s' utils.PathIsOk returned the error: %s", path, err.Error())
  }
}

func TestPathIsOk_withDoubleSlashesStart(t *testing.T) {
  path := "//a/b/c/d/"
  err := utils.PathIsOk(path)
  if err == nil {
    t.Fatalf("For '%s' utils.PathIsOk returned no error", path)
  }
}

func TestPathIsOk_withDoubleSlashesEnd(t *testing.T) {
  path := "/a/b/c/d//"
  err := utils.PathIsOk(path)
  if err == nil {
    t.Fatalf("For '%s' utils.PathIsOk returned no error", path)
  }
}

func TestPathIsOk_withDoubleSlashesMiddle(t *testing.T) {
  path := "/a/b//c/d/"
  err := utils.PathIsOk(path)
  if err == nil {
    t.Fatalf("For '%s' utils.PathIsOk returned no error", path)
  }
}

func TestPathIsOk_withDoubleSlashesDoubleDot(t *testing.T) {
  path := "/a/b/../c/d/"
  err := utils.PathIsOk(path)
  if err == nil {
    t.Fatalf("For '%s' utils.PathIsOk returned no error", path)
  }
}

func TestPathIsOk_withDoubleSlashesBackSlash(t *testing.T) {
  path := "/a/b/\\/c/d/"
  err := utils.PathIsOk(path)
  if err == nil {
    t.Fatalf("For '%s' utils.PathIsOk returned no error", path)
  }
}
