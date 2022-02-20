package usecase_test

import (
	"testing"

	mock "ratri/src/mock/repository"
	"ratri/src/usecase"

	"github.com/golang/mock/gomock"
)

func TestFetchSerp(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	sm := mock.NewMockSerpRepository(c)
	lm := mock.NewMockLinkedPageRepository(c)

	usecase.NewSerpUsecase(sm, lm)
}

func TestFetchSerpWithIcon(t *testing.T) {}

func TestFetchSerpWithRatio(t *testing.T) {}
