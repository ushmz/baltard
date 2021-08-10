package dao

import (
	"fmt"
	"testing"

	"baltard/api/models"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	// [TODO] Insert test data
	tests := []struct {
		userId   int
		expected models.ExistUser
	}{
		// {userId: 150, expected: models.ExistUser{}},
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
		expected models.ExistUser
	}{
		// {uid: "", expected: models.ExistUser{}},
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

func TestInsertCompletionCode(t *testing.T) {
	// [TODO] Insert test data
	tests := []struct {
		userId   int
		code     int
		expected error
	}{
		// {userId: 150, code: 42, expected: nil},
	}

	for idx, test := range tests {
		err := userDao.InsertCompletionCode(test.userId, test.code)

		assert.Equal(
			t,
			test.expected,
			err,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}

}

func TestGetCompletionCodeById(t *testing.T) {
	// [TODO] Insert test data
	tests := []struct {
		userId   int
		expected models.ExistUser
	}{
		// {userId: 150, expected: models.ExistUser{}},
	}

	for idx, test := range tests {
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
