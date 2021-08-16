package dao

import (
	"fmt"
	"testing"

	"baltard/api/model"

	"github.com/stretchr/testify/assert"
)

var (
	expected5 = &model.Task{
		Id:          5,
		Query:       "ウェブカメラ おすすめ",
		Title:       "購入するウェブカメラのメーカー探し",
		Description: "あなたは今「リモートでのやり取りが増えたため、ウェブカメラを購入しようと考えており、ウェブ検索をして情報を集めようとしている」とします。このページの下にある「検索結果リストを表示する」ボタンをクリックしてウェブ検索を開始してください。表示されたリストに含まれる情報を参考にして、購入したいウェブカメラのメーカーを1つ決めてください。メーカーが決まったらウェブ検索を終了し、このページの末尾にあるタスク回答欄にあなたの回答を入力して下さい。その際、回答の理由も添えてください。回答が終了したら、ページ末尾の「回答を提出する」ボタンをクリックし、次のタスクに進んでください。",
		SearchUrl:   "webcam",
	}
	expected6 = &model.Task{
		Id:          6,
		Query:       "糖尿病 症状",
		Title:       "糖尿病の症状",
		Description: "あなたは今「『糖尿病』について知りたいと思い、ウェブ検索をして調べようとしている」とします。このページの下にある「検索結果リストを表示する」ボタンをクリックしてウェブ検索を開始してください。表示されたリストに含まれる情報を参考にして、メニエール病とはどのような症状があるのかを調べてください。納得する結論が得られたらウェブ検索を終了し、このページの末尾にあるタスク回答欄にあなたの回答を入力して下さい。その際、回答の理由も添えてください。回答が終了したら、ページ末尾の「回答を提出する」ボタンをクリックし、次のタスクに進んでください。",
		SearchUrl:   "diabetes",
	}
	expected7 = &model.Task{
		Id:          7,
		Query:       "イヤホン おすすめ",
		Title:       "購入するイヤホンのメーカー探し",
		Description: "あなたは今「イヤホンを購入しようと考えており、ウェブ検索をして情報を集めようとしている」とします。このページの下にある「検索結果リストを表示する」ボタンをクリックしてウェブ検索を開始してください。表示されたリストに含まれる情報を参考にして、購入したいイヤホンのメーカーを1つ決めてください。メーカーが決まったらウェブ検索を終了し、このページの末尾にあるタスク回答欄にあなたの回答を入力して下さい。その際，回答の理由も添えてください。回答が終了したら、ページ末尾の「回答を提出する」ボタンをクリックし、次のタスクに進んでください。",
		SearchUrl:   "earphone",
	}
	expected8 = &model.Task{
		Id:          8,
		Query:       "メニエール病 症状",
		Title:       "メニエール病の症状",
		Description: "あなたは今「『メニエール病』について知りたいと思い、ウェブ検索をして調べようとしている」とします。このページの下にある「検索結果リストを表示する」ボタンをクリックしてウェブ検索を開始してください。表示されたリストに含まれる情報を参考にして、メニエール病とはどのような症状があるのかを調べてください。納得する結論が得られたらウェブ検索を終了し、このページの末尾にあるタスク回答欄にあなたの回答を入力して下さい。その際、回答の理由も添えてください。回答が終了したら、ページ末尾の「回答を提出する」ボタンをクリックし、次のタスクに進んでください。",
		SearchUrl:   "meniere",
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
		expected *model.Task
	}{

		{taskId: 5, expected: expected5},
		{taskId: 6, expected: expected6},
		{taskId: 7, expected: expected7},
		{taskId: 8, expected: expected8},
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
