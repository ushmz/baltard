package mysql_test

import (
	"testing"
)

func TestLinkedPageGet(t *testing.T) {
	if _, err := linkedPageDao.Get(1); err != nil {
		t.Fatal(err)
	}
}

func TestGetBySearchPageId(t *testing.T) {
	if _, err := linkedPageDao.GetBySearchPageIDs([]int{381, 382, 383, 384, 385}, 5, 10); err != nil {
		t.Fatal(err)
	}
}

func TestGetRatioBySearchPageId(t *testing.T) {
	if _, err := linkedPageDao.GetRatioBySearchPageIDs([]int{381, 382, 383, 384, 385}, 5); err != nil {
		t.Fatal(err)
	}
}

func TestLinkedPageSelect(t *testing.T) {
	if _, err := linkedPageDao.Select([]int{1, 3, 5}); err != nil {
		t.Fatal(err)
	}
}

func TestLinkedPageList(t *testing.T) {
	if _, err := linkedPageDao.List(0, 10); err != nil {
		t.Fatal(err)
	}
}
