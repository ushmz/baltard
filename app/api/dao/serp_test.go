package dao

import (
	"fmt"
	"testing"

	"baltard/api/model"

	"github.com/stretchr/testify/assert"
)

func TestFetchSerpByID(t *testing.T) {
	tests := []struct {
		taskId   int
		offset   int
		expected []model.SearchPage
	}{
		// [TODO] Write test case
		// {taskId: 5, offset: 0, expected: []model.SearchPage{}},
		// {taskId: 5, offset: 2, expected: []model.SearchPage{}},
		// {taskId: 6, offset: 4, expected: []model.SearchPage{}},
		// {taskId: 6, offset: 6, expected: []model.SearchPage{}},
		// {taskId: 7, offset: 8, expected: []model.SearchPage{}},
		// {taskId: 7, offset: 10, expected: []model.SearchPage{}},
		// {taskId: 8, offset: 12, expected: []model.SearchPage{}},
		// {taskId: 8, offset: 14, expected: []model.SearchPage{}},
	}

	for idx, test := range tests {
		srp, err := serpDao.FetchSerpByID(test.taskId, test.offset)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			srp,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}

func TestFetchSerpWithIconByID(t *testing.T) {
	tests := []struct {
		taskId   int
		offset   int
		top      int
		expected []model.SearchPage
	}{
		// [TODO] Write test case
		// {taskId: 5, offset: 0, top: 10, expected: []model.SearchPage{}},
		// {taskId: 5, offset: 2, top: 10, expected: []model.SearchPage{}},
		// {taskId: 6, offset: 4, top: 10, expected: []model.SearchPage{}},
		// {taskId: 6, offset: 6, top: 10, expected: []model.SearchPage{}},
		// {taskId: 7, offset: 8, top: 10, expected: []model.SearchPage{}},
		// {taskId: 7, offset: 10, top: 0, expected: []model.SearchPage{}},
		// {taskId: 8, offset: 12, top: 0, expected: []model.SearchPage{}},
		// {taskId: 8, offset: 14, top: 0, expected: []model.SearchPage{}},
	}

	for idx, test := range tests {
		srp, err := serpDao.FetchSerpWithIconByID(test.taskId, test.offset, test.top)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			srp,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}
func TestFetchSerpWithDistributionByID(t *testing.T) {
	tests := []struct {
		taskId   int
		offset   int
		top      int
		expected []model.SearchPage
	}{
		// [TODO] Write test case
		// {taskId: 5, offset: 0, top: 10, expected: []model.SearchPage{}},
		// {taskId: 5, offset: 2, top: 10, expected: []model.SearchPage{}},
		// {taskId: 6, offset: 4, top: 10, expected: []model.SearchPage{}},
		// {taskId: 6, offset: 6, top: 10, expected: []model.SearchPage{}},
		// {taskId: 7, offset: 8, top: 10, expected: []model.SearchPage{}},
		// {taskId: 7, offset: 10, top: 0, expected: []model.SearchPage{}},
		// {taskId: 8, offset: 12, top: 0, expected: []model.SearchPage{}},
		// {taskId: 8, offset: 14, top: 0, expected: []model.SearchPage{}},
	}

	for idx, test := range tests {
		srp, err := serpDao.FetchSerpWithDistributionByID(test.taskId, test.offset, test.top)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			srp,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}
