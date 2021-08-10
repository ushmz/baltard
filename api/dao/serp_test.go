package dao

import (
	"fmt"
	"testing"

	"baltard/api/models"

	"github.com/stretchr/testify/assert"
)

func TestFetchSerpByID(t *testing.T) {
	tests := []struct {
		taskId   int
		offset   int
		expected []models.SearchPage
	}{
		// [TODO] Write test case
		// {taskId: 5, offset: 0, expected: []models.SearchPage{}},
		// {taskId: 5, offset: 2, expected: []models.SearchPage{}},
		// {taskId: 6, offset: 4, expected: []models.SearchPage{}},
		// {taskId: 6, offset: 6, expected: []models.SearchPage{}},
		// {taskId: 7, offset: 8, expected: []models.SearchPage{}},
		// {taskId: 7, offset: 10, expected: []models.SearchPage{}},
		// {taskId: 8, offset: 12, expected: []models.SearchPage{}},
		// {taskId: 8, offset: 14, expected: []models.SearchPage{}},
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
		expected []models.SearchPage
	}{
		// [TODO] Write test case
		// {taskId: 5, offset: 0, top: 10, expected: []models.SearchPage{}},
		// {taskId: 5, offset: 2, top: 10, expected: []models.SearchPage{}},
		// {taskId: 6, offset: 4, top: 10, expected: []models.SearchPage{}},
		// {taskId: 6, offset: 6, top: 10, expected: []models.SearchPage{}},
		// {taskId: 7, offset: 8, top: 10, expected: []models.SearchPage{}},
		// {taskId: 7, offset: 10, top: 0, expected: []models.SearchPage{}},
		// {taskId: 8, offset: 12, top: 0, expected: []models.SearchPage{}},
		// {taskId: 8, offset: 14, top: 0, expected: []models.SearchPage{}},
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
		expected []models.SearchPage
	}{
		// [TODO] Write test case
		// {taskId: 5, offset: 0, top: 10, expected: []models.SearchPage{}},
		// {taskId: 5, offset: 2, top: 10, expected: []models.SearchPage{}},
		// {taskId: 6, offset: 4, top: 10, expected: []models.SearchPage{}},
		// {taskId: 6, offset: 6, top: 10, expected: []models.SearchPage{}},
		// {taskId: 7, offset: 8, top: 10, expected: []models.SearchPage{}},
		// {taskId: 7, offset: 10, top: 0, expected: []models.SearchPage{}},
		// {taskId: 8, offset: 12, top: 0, expected: []models.SearchPage{}},
		// {taskId: 8, offset: 14, top: 0, expected: []models.SearchPage{}},
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
