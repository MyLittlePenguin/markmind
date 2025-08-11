package utils

import (
	"errors"
	"strings"
)

func Ternary[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}
	return b
}

func PathIsOk(path string) error {
	if strings.Contains(path, "..") {
		return errors.New("'..' are not allowed inside the path")
	}
	if strings.Contains(path, "//") {
		return errors.New("'//' are not allowed inside the path")
	}
  if strings.Contains(path, "\\") {
		return errors.New("'\\' are not allowed inside the path")
  }
  return nil
}

func SanitizePath(path string) (string, error) {
  err := PathIsOk(path)
  if err != nil {
    return "", err
  }

  originalPathSegments := strings.Split(path, "/")
  finalPathSegments := make([]string,0 , 10)
  for i := 0; i < len(originalPathSegments); i++ {
    segment := originalPathSegments[i]
    if segment != "" {
      finalPathSegments = append(finalPathSegments, segment)
    }
  }

  return strings.Join(finalPathSegments, "/"), nil
}
