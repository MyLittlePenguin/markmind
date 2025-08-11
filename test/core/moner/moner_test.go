package moner

import (
	"errors"
	"markmind/internal/core/moner"
	"strconv"
	"testing"
)

func moaner(num int) (string, error) {
	switch num {
	case 0:
		return "42", nil
	case 1:
		return "q", nil
	default:
		return "", errors.New("Möp")
	}
}

func TestBind(t *testing.T) {
	resultMoner := moner.Bind(
		moaner,
		strconv.Atoi,
	)

	v, err := resultMoner(0)
	if err != nil {
		t.Fatalf("err is unexpectedly nil, v has value: %d", v)
	}
}

func TestCompose(t *testing.T) {

	resultMoner := moner.Compose(
		moner.WrapFn(moaner),
		moner.WrapFn(strconv.Atoi),
	)
  
	v, err := resultMoner(0)
	if err != nil || v != 42 {
		t.Fatalf("unexpectedly result, (v, err): (%d, %s)", v, err)
	}

	v, err = resultMoner(1)
	if err == nil || v != 0 {
		t.Fatalf("unexpectedly result, (v, err): (%d, %s)", v, err)
	}

	v, err = resultMoner(2)
	if err == nil {
		t.Fatalf("unexpectedly result, (v, err): (%s, %s)", v, err)
	}
}
