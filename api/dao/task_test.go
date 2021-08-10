package dao

import (
	"fmt"
	"testing"

	"baltard/api/models"

	"github.com/stretchr/testify/assert"
)

var (
	expTask5 = &models.Task{
		Id:          5,
		Query:       "ウェブカメラ おすすめ",
		Title:       "購入するウェブカメラのメーカー探し",
		Description: `あなたは今「リモートでのやり取りが増えたため、ウェブカメラを購入しようと考えており、ウェブ検索をして情報を集めようとしている」とします。このページの下にある「検索結果リストを表示する」ボタンをクリックしてウェブ検索を開始してください。表示されたリストに含まれる情報を参考にして、購入したいウェブカメラのメーカーを1つ決めてください。メーカーが決まったらウェブ検索を終了し、このページの末尾にあるタスク回答欄にあなたの回答を入力して下さい。その際、回答の理由も添えてください。回答が終了したら，ページ末尾の「回答を提出する」ボタンをクリックし、次のタスクに進んでください。`,
		SearchUrl:   "webcam",
	}
)

// [TODO] How to test this?
func TestAllocateTask(t *testing.T) {}

func TestFetchTaskIdsByGroupId(t *testing.T) {
	tests := []struct {
		groupId  int
		expected []int
	}{
		{groupId: 1, expected: []int{5, 7}},
		{groupId: 2, expected: []int{6, 8}},
		{groupId: 3, expected: []int{5, 7}},
		{groupId: 4, expected: []int{6, 8}},
		{groupId: 5, expected: []int{5, 7}},
		{groupId: 6, expected: []int{6, 8}},
	}

	for idx, test := range tests {
		ids, err := taskDao.FetchTaskIdsByGroupId(test.groupId)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			ids,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}

}

func TestFetchTaskInfo(t *testing.T) {
	tests := []struct {
		taskId   int
		expected *models.Task
	}{

		{taskId: 5, expected: expTask5},
	}

	for idx, test := range tests {
		tsk, err := taskDao.FetchTaskInfo(test.taskId)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(
			t,
			test.expected,
			tsk,
			fmt.Sprintf("Testcase index %d", idx),
		)
	}
}
