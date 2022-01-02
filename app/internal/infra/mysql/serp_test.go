package mysql_test

import "testing"

func TestFetchSerpByTaskID(t *testing.T) {
	if _, err := serpDao.FetchSerpByTaskID(5, 0); err != nil {
		t.Fatal(err)
	}
}
func TestFetchSerpWithIconByTaskID(t *testing.T) {
	if _, err := serpDao.FetchSerpWithIconByTaskID(5, 0, 10); err != nil {
		t.Fatal(err)
	}
}
func TestFetchSerpWithRatioByTaskID(t *testing.T) {
	if _, err := serpDao.FetchSerpWithRatioByTaskID(5, 0, 3); err != nil {
		t.Fatal(err)
	}
}
