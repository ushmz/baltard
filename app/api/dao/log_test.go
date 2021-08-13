package dao

import (
	"fmt"
	"testing"

	"baltard/api/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateTaskTimeLog(t *testing.T) {
	tests := []struct {
		log      models.TaskTimeLogParam
		expected error
	}{
		{
			log: models.TaskTimeLogParam{
				UserId:      999,
				TimeOnPage:  999,
				TaskId:      5,
				ConditionId: 5,
			},
			expected: nil,
		},
		{
			log: models.TaskTimeLogParam{
				UserId:      999,
				TimeOnPage:  -999,
				TaskId:      5,
				ConditionId: 5,
			},
			expected: nil,
		},

		{
			log: models.TaskTimeLogParam{
				UserId:      999,
				TimeOnPage:  999,
				TaskId:      -5,
				ConditionId: 5,
			},
			expected: nil,
		},
		{
			log: models.TaskTimeLogParam{
				UserId:      999,
				TimeOnPage:  999,
				TaskId:      5,
				ConditionId: -5,
			},
			expected: nil,
		},
	}

	for idx, test := range tests {
		err := logDao.CreateTaskTimeLog(&test.log)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(
			t,
			test.expected,
			err,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}
