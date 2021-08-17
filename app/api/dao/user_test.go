package dao

import (
	"fmt"
	"testing"

	"baltard/api/model"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	// [TODO] Insert test data
	tests := []struct {
		userId   int
		expected model.User
	}{
		// {userId: 150, expected: model.ExistUser{}},
	}

	for idx, test := range tests {
		eu, err := userDao.FindById(test.userId)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			eu,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}

func TestFindByUid(t *testing.T) {
	// [TODO] Insert test data
	tests := []struct {
		uid      string
		expected model.User
	}{
		// {uid: "", expected: model.ExistUser{}},
	}

	for idx, test := range tests {
		eu, err := userDao.FindByUid(test.uid)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			eu,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}

func TestCompletionCodeIO(t *testing.T) {
	tests := []struct {
		userId   int
		code     int
		expected int
	}{
		{userId: 9999, code: 42, expected: 42},
		{userId: 99999, code: 4242, expected: 4242},
		{userId: 999999, code: 4422, expected: 4422},
		{userId: 444422, code: 44422, expected: 44422},
	}

	for idx, test := range tests {
		err := userDao.InsertCompletionCode(test.userId, test.code)
		if err != nil {
			t.Fatal(err)
		}

		code, err := userDao.GetCompletionCodeById(test.userId)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			code,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}

}
