package handler

import (
	"testing"
)

var (
	taskInfo = `[
    	{
        	"id": 5,
			"query": "ウェブカメラ おすすめ",
			"title": "購入するウェブカメラのメーカー探し",
			"description": "あなたは今「リモートでのやり取りが増えたため、ウェブカメラを購入しようと考えており、ウェブ検索をして情報を集めようとしている」とします。このページの下にある「検索結果リストを表示する」ボタンをクリックしてウェブ検索を開始してください。表示されたリストに含まれる情報を参考にして、購入したいウェブカメラのメーカーを1つ決めてください。メーカーが決まったらウェブ検索を終了し、このページの末尾にあるタスク回答欄にあなたの回答を入力して下さい。その際、回答の理由も添えてください。回答が終了したら，ページ末尾の「回答を提出する」ボタンをクリックし、次のタスクに進んでください。",
			"authorId": "",
			"searchUrl": "webcam",
			"type": {
				"String": "",
				"Valid": false
			}
		}
	]`
	answer = `{
		"uid": "test_uid",
		"task": 5,
		"condition": 5,
		"authorId": 2,
		"answer": "test_answer",
		"reason": "test_reason"
	}`
)

func TestFetchTaskInfo(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/v1/task/5", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.FetchTaskInfo(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusOK); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// 	if diff := cmp.Diff(rec.Body.String(), taskInfo); diff != "" {
	// 		t.Errorf("Response body does not match.\n%v", diff)
	// 	}
	// }
}

func TestSubmitTaskAnswer(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/v1/task/answer", strings.NewReader(answer))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.SubmitTaskAnswer(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusCreated); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}
