package mysql_test

import "testing"

func TestFetchSerpByTaskID(t *testing.T) {
	if _, err := serpDao.FetchSerpByTaskID(5, 0); err != nil {
		t.Fatal(err)
	}
}
