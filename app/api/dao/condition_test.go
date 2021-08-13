package dao

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchConditionIdByGroupId(t *testing.T) {
	tests := []struct {
		groupId  int
		expected int
	}{
		{groupId: 0, expected: 0},
		{groupId: 1, expected: 5},
		{groupId: 2, expected: 5},
		{groupId: 3, expected: 6},
		{groupId: 4, expected: 6},
		{groupId: 5, expected: 7},
		{groupId: 6, expected: 7},
		{groupId: 7, expected: 0},
	}

	for idx, test := range tests {
		cond, err := conditionDao.FetchConditionIdByGroupId(test.groupId)
		if err != nil {
			if err == sql.ErrNoRows {
				assert.Equal(
					t,
					test.expected,
					cond,
					fmt.Sprintf("Testcase index %d", idx),
				)
			} else {

				t.Fatal(err)
			}
		}

		assert.Equal(
			t,
			test.expected,
			cond,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}
