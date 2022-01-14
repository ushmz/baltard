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
	if v, err := linkedPageDao.GetBySearchPageIds([]int{381, 382, 383, 384, 385}, 5, 10); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%#v", v)
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
